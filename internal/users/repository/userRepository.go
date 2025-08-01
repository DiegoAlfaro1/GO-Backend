package repository

import (
	"database/sql"
	"fmt"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetOneUser(userID string) (model.User, error)
	GetOneUserByEmail(email string) (model.User, error)
	CreateFromCognito(user model.User) (model.User, error)
	DeleteOne(userID string) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{db: config.DB}
}

func (r *userRepo) GetAll() ([]model.User, error) {
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

func (r *userRepo) GetOneUser(userID string) (model.User, error) {
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

func (r *userRepo) GetOneUserByEmail(email string) (model.User, error) {
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

func (r *userRepo) CreateFromCognito(user model.User) (model.User, error) {
	_, err := r.db.Exec(
		"INSERT INTO users (id, name, email, birthdate) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING",
		user.ID, user.Name, user.Email, user.Birthdate,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepo) DeleteOne(userID string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}
