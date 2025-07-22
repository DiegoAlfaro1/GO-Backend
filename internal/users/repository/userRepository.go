package repository

import (
	"context"
	"time"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface{
	GetAll() ([]model.User, error)
	GetOneUser(userID string) (model.User, error)
	Create(user model.User) (model.User, error)
	DeleteOne(userID string) (error)
}

type userRepo struct{
	collection *mongo.Collection
}

func NewUserRepository() UserRepository {
	col := config.GetCollection("users")
    return &userRepo{collection: col}
}

func (r *userRepo) GetAll() ([]model.User, error){
	var users []model.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx){
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users,user)
	}

	return users, nil
}

func (r *userRepo) GetOneUser(userID string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}


func (r *userRepo) Create(user model.User) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	return user, err
}

func (r *userRepo) DeleteOne(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
