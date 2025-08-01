package repository

import (
	"database/sql"
	"fmt"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/model"
	"github.com/DiegoAlfaro1/gin-terraform/internal/util"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetOneUser(userID string) (model.User, error)
	GetOneUserByEmail(email string) (model.User, error)
	CreateFromCognito(email string) (error)
	DeleteOne(userID string) error
}

type UserRepo struct {
	db *sql.DB
	cognitoClient config.CognitoInterface
}

func NewUserRepository(cognitoClient config.CognitoInterface) UserRepository {
	return &UserRepo{
		db: config.DB,
		cognitoClient: cognitoClient,
	
	}
}

func (r *UserRepo) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, COALESCE(birthdate, '') FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Birthdate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepo) GetOneUser(userID string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(
		"SELECT id, name, email, COALESCE(birthdate, '') FROM users WHERE id = $1",
		userID,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Birthdate)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetOneUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(
		"SELECT id, name, email, COALESCE(birthdate, '') FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Birthdate)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) CreateFromCognito(email string) (error) {

	rawUser, userErr := r.cognitoClient.GetUserFromEmail(email)

	if userErr != nil {
		return userErr
	}

	user := util.ExtractAttributes(rawUser)

	_, err := r.db.Exec(
		"INSERT INTO users (id, name, email, birthdate) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING",
		user["custom:custom_id"], user["name"], user["email"], user["birthdate"],
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteOne(userID string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}
