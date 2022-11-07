package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCongig struct {
	HOST string
	PORT string
}

func NewMongoDB(cfg MongoCongig, ctx context.Context) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s", cfg.HOST, cfg.PORT)

	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	mdb, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	if err = mdb.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return mdb, nil
}
