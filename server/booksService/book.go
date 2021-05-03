package library

import "context"

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category,omitempty"`
}

type Repository interface {
	getBooks(ctx context.Context) (string, error)
	addBook(ctx context.Context, b Book) (string, error)
}
