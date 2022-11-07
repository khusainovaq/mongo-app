package repository

import (
	"mongo-l3/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type Auth interface {
	CreateUser(u entity.User) error
	GetUser(username interface{}) (entity.User, error)
	EditUser(username interface{}, u entity.User) error
	DeleteUser(username interface{}) error
}

type Repository struct {
	Auth
}

func NewRepository(mdb *mongo.Client) *Repository {
	return &Repository{
		Auth: NewAuthMongo(mdb),
	}
}
