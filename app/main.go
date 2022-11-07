package main

import (
	"context"
	"os"

	code "mongo-l3"
	"mongo-l3/driver"
	"mongo-l3/pkg/hanlder"
	"mongo-l3/pkg/repository"
	"mongo-l3/pkg/service"

	"github.com/joho/godotenv"
	"github.com/khusainnov/logging"
)

var (
	logger = logging.GetLogger()
	ctx    = context.Background()
)

func main() {
	if err := godotenv.Load("./config/.env"); err != nil {
		logger.Errorf("Cannot load config, due to error: %s", err.Error())
	}

	mdb, err := driver.NewMongoDB(driver.MongoCongig{
		HOST: os.Getenv("MONGO_HOST"),
		PORT: os.Getenv("MONGO_PORT"),
	},
		ctx,
	)
	if err != nil {
		logger.Fatalf("Error due connect to mongo: %s", err.Error())
	}

	repo := repository.NewRepository(mdb)
	services := service.NewService(repo)
	h := hanlder.NewHandler(services)
	s := code.Server{}

	logger.Infof("Starting server on port:%s", os.Getenv("PORT"))
	if err = s.RunHTTP(os.Getenv("PORT"), h.InitRoutes()); err != nil {
		logger.Fatalf("Error due start the server: %s", err.Error())
	}
}
