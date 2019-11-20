package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		return nil, err
	}
	return loginRequest, nil
}

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var registerRequest RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		return nil, err
	}
	return registerRequest, nil
}

func encodeRegisterResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func MakeHttpHandler(ctx context.Context, endpoint AuthEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	r.Methods("POST").Path("/api/v1/auth/sessions").Handler(kithttp.NewServer(
		endpoint.LoginEndpoint,
		decodeLoginRequest,
		encodeLoginResponse,
		options...,
	))

	r.Methods("POST").Path("/api/v1/auth/users").Handler(kithttp.NewServer(
		endpoint.RegisterEndpoint,
		decodeRegisterRequest,
		encodeLoginResponse,
		options...,
	))
	return r

}
