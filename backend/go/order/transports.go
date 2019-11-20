package order

import (
	"context"
	"encoding/json"
	"net/http"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeAllRequest(c context.Context, r *http.Request) (interface{}, error) {
	var allOrderRequest AllOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&allOrderRequest); err != nil {
		return nil, err
	}
	return allOrderRequest, nil
}

func encodeAllResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func MakeHttpHandler(ctx context.Context, endpoint OrderEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
	}

	r.Methods("GET").Path("/api/v1/order/all").Handler(kithttp.NewServer(
		endpoint.AllOrderEndpoint,
		decodeAllRequest,
		encodeAllResponse,
		options...,
	))

	return r
}
