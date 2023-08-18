package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"umbrella.github.com/go-grpc-example/server/help"
	"umbrella.github.com/go-grpc-example/server/services"
)

func main() {

	//ssl
	// creds, err := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	grpc_server := grpc.NewServer(grpc.Creds(help.GetServerCreds()))
	services.RegisterProdServiceServer(grpc_server, new(services.ProdService))
	services.RegisterOrderServiceServer(grpc_server, new(services.OrdersService))
	services.RegisterUserServiceServer(grpc_server, new(services.UserService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {

		log.Fatal(err)
	}

	grpc_server.Serve(listen)
}
