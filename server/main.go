package main

import (
	"context"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(
			logger,
			"service", "library",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	Level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	clientDBConn, err := mongo.NewClient(options.Client().ApplyURI("localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = clientDBConn.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer clientDBConn.Disconnect(ctx)

	libraryDb := clientDBConn.Database("books_db")
	booksCol := libraryDb.Collection("books")

	var srv library.Service
	{
		repository := library.NewRepo(booksCol, logger)
		srv = library.NewService(repository, logger)
	}

	endpoints := library.MakeEndpoints(srv)
}
