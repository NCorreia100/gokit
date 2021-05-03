package library

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct {
	db     *mongodb.DB
	logger log.Logger
}

var RepoErr = errors.New("Unable to handle Repo request")

func NewRepo(db *mongodb.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongoDb"),
	}
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	libraryDb := client.Database("books_db")
	booksCol := libraryDb.Collection("books")
}
