package service

import (
	"fmt"

	"github.com/DiegoAlfaro1/gin-terraform/internal/users/model"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/repository"
	"github.com/DiegoAlfaro1/gin-terraform/internal/util"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetOneUser(userID string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	DeleteOne(userID string) (error)

}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo}
}

func (s *userServiceImpl) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userServiceImpl) GetOneUser(userID string) (model.User, error) {
	return s.repo.GetOneUser(userID)
}
 
func (s *userServiceImpl) CreateUser(user model.User) (model.User, error) {

	hashedPassword, err := util.HashPassword(user.Password)

	if err != nil {
		return model.User{}, fmt.Errorf("error hashing password: %w", err)
	}

	user.Password = hashedPassword

	
	return s.repo.Create(user)
}

func (s *userServiceImpl) DeleteOne(userID string) (error) {
	return s.repo.DeleteOne(userID)
}