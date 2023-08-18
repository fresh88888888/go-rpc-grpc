package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"umbrella.github.com/go-grpc-example/server/help"
	"umbrella.github.com/go-grpc-example/server/services"
)

func main() {
	gmux := runtime.NewServeMux()
	grpc_endpoint := "localhost:8081"
	opt := []grpc.DialOption{grpc.WithTransportCredentials(help.GetClientCreds())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gmux, grpc_endpoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gmux, grpc_endpoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8082", gmux)
}
