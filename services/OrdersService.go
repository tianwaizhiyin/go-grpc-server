package services

import (
	"context"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type OrdersService struct {

}

func (this *OrdersService) NewOrder(ctx context.Context, orderRequest *OrderRequest) (*OrderResponse, error)  {
	err := orderRequest.OrderMain.Validate()
	if err!=nil {
		return &OrderResponse{
			state:         protoimpl.MessageState{},
			sizeCache:     0,
			unknownFields: nil,
			Status:        "error",
			Message:       err.Error(),
		}, nil
	}
	//fmt.Println(orderRequest.OrderMain)
	return &OrderResponse{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Status:        "OK",
		Message:       "sucess",
	}, nil
}