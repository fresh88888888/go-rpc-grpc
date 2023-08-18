package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"umbrella.github.com/go-grpc-example/client/help"
	"umbrella.github.com/go-grpc-example/client/services"
)

func main() {

	// creds, err := credentials.NewClientTLSFromFile("keys/server.pem", "www.umbrella.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(help.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	pro_client := services.NewProdServiceClient(conn)
	// pro_repose, err := pro_client.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12, ProdArea: services.ProdAreas_A})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(pro_repose.ProdStock)

	// ctx := context.Background()

	// response, err := pro_client.GetProdStocks(ctx, &services.QuerySize{Size: 10})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(response.Prores)
	pro_info, err := pro_client.GetProdInfo(context.Background(), &services.ProdRequest{ProdId: 12, ProdArea: services.ProdAreas_A})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pro_info)

	ord_client := services.NewOrderServiceClient(conn)
	order_info, err := ord_client.NewOrder(context.Background(), &services.OrderRequest{
		OrderMain: &services.OrderMain{
			OrderId:    1000,
			OrderNo:    "o122333",
			UserId:     101,
			OrderMoney: 12.34,
			OrderTime:  &timestamppb.Timestamp{Seconds: time.Now().Unix()},
		},
	})

	fmt.Println(order_info)

	user_client := services.NewUserServiceClient(conn)
	req := services.UserScoreRequest{}
	req.Users = make([]*services.UserInfo, 0)
	for i := 0; i < 20; i++ {
		req.Users = append(req.Users, &services.UserInfo{UserId: int32(i), UserScore: int32(i) + 12})
	}

	stream, err := user_client.GetUserScoreByServerStream(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response.Users)
	}
}
