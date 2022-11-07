package service

import (
	"mongo-l3/internal/entity"
	"mongo-l3/pkg/repository"
)

type Auth interface {
	CreateUser(u entity.User) error
	GetUser(username interface{}) (entity.User, error)
	EditUser(username interface{}, u entity.User) error
	DeleteUser(username interface{}) error
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
