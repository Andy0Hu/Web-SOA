package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type OrderEndpoints struct {
	AllOrderEndpoint endpoint.Endpoint
}

type AllOrderRequest struct {
	Id string `json:"id"`
}

type AllOrderResponse struct {
	All   []Orders `json:"all"`
	Error string   `json:"error"`
}

func MakeAllOrderEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AllOrderRequest)

		all, err := svc.AllOrders(req.Id)
		var resp AllOrderResponse
		if err != nil {
			resp.Error = err.Error()
		} else {
			resp.All = all.([]Orders)
			resp.Error = ""
		}
		return resp, nil
	}
}
