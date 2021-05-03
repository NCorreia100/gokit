package library

import "context"

type Service interface {
	getBooks(ctx context.Context) (string, error)
	addBook(ctx context.Context, name string, author string, category string) (string, error)
}
