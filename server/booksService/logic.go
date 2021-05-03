package library

import (
	"context"

	"github.com/go-kit/kit/log"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) GetBooks(ctx context.Context) (string, error) {
	logger := log.With(s.logger, "method", "GetBooks")

	bs, err := s.repository.GetBooks(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get book")

	return bs, nil
}

func (s service) AddBook(ctx context.Context, title string, author string, category string) (string, error) {
	logger := log.With(s.logger, "method", "AddBook")

	book := Book{
		Title:    title,
		author:   author,
		Category: category,
	}

	if err := s.repository.AddBook(ctx, book); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Book added", book)

	return "Success", nil
}
