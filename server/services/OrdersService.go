package services

import (
	context "context"
	"fmt"
)

type OrdersService struct {
}

func (*OrdersService) mustEmbedUnimplementedOrderServiceServer() {

}

func (*OrdersService) NewOrder(ctx context.Context, order *OrderRequest) (*OrderResponse, error) {
	fmt.Println(order.OrderMain)
	return &OrderResponse{Status: "open", Messages: "this is order"}, nil
}
