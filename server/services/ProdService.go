package services

import (
	context "context"
)

type ProdService struct {
}

func (*ProdService) GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error) {
	pro_info := &ProdModel{
		ProdId:    101,
		ProdName:  "tomatosa",
		ProdPrice: 12.23,
	}

	return pro_info, nil
}

func (*ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	prodres := []*ProdResponse{
		{ProdStock: 20},
		{ProdStock: 21},
		{ProdStock: 22},
		{ProdStock: 23},
		{ProdStock: 24},
		{ProdStock: 25},
		{ProdStock: 26},
	}

	return &ProdResponseList{Prores: prodres}, nil
}

func (this *ProdService) mustEmbedUnimplementedProdServiceServer() {

}

func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	if req.ProdArea == ProdAreas_A {
		stock = 30
	} else if req.ProdArea == ProdAreas_B {
		stock = 31
	} else {
		stock = 50
	}
	return &ProdResponse{ProdStock: stock}, nil
}
