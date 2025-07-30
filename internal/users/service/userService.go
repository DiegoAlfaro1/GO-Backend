package service

import (
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/model"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/repository"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetOneUser(userID string) (model.User, error)
	CreateUserFromCognito(name, email, customID string) (model.User, error)
	CreateUserFromEmail(email string) (model.User, error)
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
 
func (s *userServiceImpl) CreateUserFromCognito(name, email, customID string) (model.User, error) {
	user := model.User{
		ID:    customID,
		Name:  name,
		Email: email,
	}
	return s.repo.CreateFromCognito(user)
}

func (s *userServiceImpl) CreateUserFromEmail(email string) (model.User, error) {

	user := model.User{
		ID:    uuid.NewString(),  // TODO: Get this data from cognito
		Email: email,
		Name:  "", // TODO: Get this data from cognito
	}
	return s.repo.CreateFromCognito(user)
}

func (s *userServiceImpl) DeleteOne(userID string) error {
	return s.repo.DeleteOne(userID)
}