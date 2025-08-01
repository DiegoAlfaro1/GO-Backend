package service

import (
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/model"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetOneUser(userID string) (model.User, error)
	CreateUserFromEmail(email string) (error)
	DeleteOne(userID string) error
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
 
func (s *userServiceImpl) CreateUserFromEmail(email string) (error) {
	return s.repo.CreateFromCognito(email)
}

func (s *userServiceImpl) DeleteOne(userID string) error {
	return s.repo.DeleteOne(userID)
}
