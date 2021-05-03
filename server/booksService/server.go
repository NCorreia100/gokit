package library

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleWare)

	r.Methods("POST").Path("/addBook").Handler(httptransport.NewServer(
		endpoints.AddBook,
		decodeRequest,
		encodeResponse
		))

	r.Methods("GET").Path("/getBooks").Handler(httptransport.NewServer(
		endpoints.GetBooks,
		decodeRequest,
		encodeResponse
		))

}

//middleware for verifying requests
func commonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "Application/json")
		next.ServeHTTP(w, r)
	})
}
