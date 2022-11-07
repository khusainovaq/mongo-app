package service

import (
	"mongo-l3/internal/entity"
	"mongo-l3/pkg/repository"

	"github.com/google/uuid"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(u entity.User) error {
	u.ID = uuid.New().String()
	if err := as.repo.CreateUser(u); err != nil {
		return err
	}
	return nil
}
func (as *AuthService) GetUser(username interface{}) (entity.User, error) {
	u, err := as.repo.GetUser(username)
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}
func (as *AuthService) EditUser(username interface{}, u entity.User) error {
	if err := as.repo.EditUser(username, u); err != nil {
		return err
	}
	return nil
}
func (as *AuthService) DeleteUser(username interface{}) error {
	if err := as.repo.DeleteUser(username); err != nil {
		return err
	}
	return nil
}
