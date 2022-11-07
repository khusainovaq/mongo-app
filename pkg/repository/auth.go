package repository

import (
	"context"

	"mongo-l3/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userDB         = "user-collection"
	userCollection = "user-data"
)

var (
	ctx = context.Background()
)

type Mongo struct {
	mdb *mongo.Client
}

func NewAuthMongo(mdb *mongo.Client) *Mongo {
	return &Mongo{mdb: mdb}
}

func (m *Mongo) CreateUser(u entity.User) error {
	uc := m.mdb.Database(userDB).Collection(userCollection)
	if _, err := uc.InsertOne(ctx, &u); err != nil {
		return err
	}

	return nil
}
func (m *Mongo) GetUser(username interface{}) (entity.User, error) {
	var u entity.User
	us := m.mdb.Database(userDB).Collection(userCollection)
	if err := us.FindOne(ctx, username).Decode(&u); err != nil {
		return entity.User{}, err
	}

	return u, nil
}
func (m *Mongo) EditUser(username interface{}, u entity.User) error {
	uc := m.mdb.Database(userDB).Collection(userCollection)

	filter := bson.M{"username": username}
	fields := bson.M{"$set": bson.M{"name": u.Name, "surname": u.Surname, "username": u.Username, "email": u.Email, "number": u.Number}}
	if err := uc.FindOneAndUpdate(ctx, filter, fields).Err(); err != nil {
		return err
	}

	return nil
}
func (m *Mongo) DeleteUser(username interface{}) error {
	uc := m.mdb.Database(userDB).Collection(userCollection)
	if _, err := uc.DeleteOne(ctx, username); err != nil {
		return err
	}

	return nil
}
