package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"soa/auth"
	"soa/order"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	svc := auth.AuthService{}
	lep := auth.MakeLoginEndpoint(svc)
	rep := auth.MakeRegisterEndpoint(svc)
	authEndpoint := auth.AuthEndpoints{
		LoginEndpoint:    lep,
		RegisterEndpoint: rep,
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := auth.MakeHttpHandler(ctx, authEndpoint, logger)

	osvc := order.OrderService{}
	ep := order.MakeAllOrderEndpoint(osvc)
	ep = kitjwt.NewParser(auth.JWTKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(ep)
	orderEndpoint := order.OrderEndpoints{
		AllOrderEndpoint: ep,
	}

	r2 := order.MakeHttpHandler(ctx, orderEndpoint, logger)

	go func() {
		fmt.Println("Http Server start at port:8080")
		handler := r
		errChan <- http.ListenAndServe(":8080", handler)
	}()

	go func() {
		fmt.Println("Http Server start at port:8081")
		handler := r2
		errChan <- http.ListenAndServe(":8081", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
