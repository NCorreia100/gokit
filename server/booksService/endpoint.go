package library

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddBook  endpoint.Endpoint
	GetBooks endpoint.Endpoint
}

func MakeEndpoints(s Service) {
	return Endpoints{
		AddBook:  makeAddBookEndpoint(s),
		GetBooks: makeGetBooksEndpoint(s),
	}
}

func makeAddBookEndpoint(s Service) endpoint.Endpoint{
	return func (ctx context.Context,request interface{}) (interface{},error)
}

func makeGetBooksEndpoint(s Service) endpoint.Endpoint{
	return func (ctx context.Context, request interface{})(interface{},error)
}