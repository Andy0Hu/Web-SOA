package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type AuthEndpoints struct {
	LoginEndpoint    endpoint.Endpoint
	RegisterEndpoint endpoint.Endpoint
}

// LoginRequest
type LoginRequest struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`
}

type RegisterRequest struct {
	Id       string `json:"id"`
	Pwd      string `json:"pwd"`
	Username string `json:"username"`
}

// LoginResponse
type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Error   string `json:"error"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func MakeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginRequest)

		token, err := svc.Login(req.Id, req.Pwd)

		var resp LoginResponse
		if err != nil {
			resp = LoginResponse{
				Success: false,
				Token:   token,
				Error:   err.Error(),
			}
		} else {
			resp = LoginResponse{
				Success: true,
				Token:   token,
			}
		}
		return resp, nil
	}
}

func MakeRegisterEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RegisterRequest)

		err = svc.Register(req.Id, req.Pwd, req.Username)

		var resp RegisterResponse
		if err != nil {
			resp = RegisterResponse{
				Success: false,
				Error:   err.Error(),
			}
		} else {
			resp = RegisterResponse{
				Success: true,
			}
		}
		return resp, nil
	}
}
