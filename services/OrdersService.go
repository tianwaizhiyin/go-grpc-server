package services

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type OrdersService struct {

}

func (this *OrdersService) NewOrder(ctx context.Context, orderMain *OrderMain) (*OrderResponse, error)  {
	fmt.Println(orderMain)
	return &OrderResponse{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Status:        "OK",
		Message:       "sucess",
	}, nil
}