package library

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetBooksRequest struct {
	}
	GetBooksResponse struct {
		//books []Book `json:struct`
	}
	AddBookRequest struct {
		Title    string `json:"title"`
		Author   string `json:"author"`
		Category string `json:"category,omitempty"`
	}
	AddBookResponse struct {
		ok string `json:"ok"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEnconder(w).Encode(response)
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req AddBookRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
