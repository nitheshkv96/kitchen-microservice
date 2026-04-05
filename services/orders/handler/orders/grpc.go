package handler

import (
	"context"
	"log"

	"github.com/nitheshkv96/kitchen-microservice/services/common/genproto/orders"
	"github.com/nitheshkv96/kitchen-microservice/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	// service Injection
	ordersService types.OrderService

	// unimplemented UnimplementedOrderServicesServer
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	// register the OrderSe5rvicesServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)

}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {

	order := &orders.Order{
		OrderID:    1,
		ProductID:  req.ProductID,
		CustomerID: req.CustonerID,
		Quantity:   req.Quantity,
	}

	err := h.ordersService.CreateOrder(ctx, order)

	if err != nil {
		log.Fatalf("Order could not be created: %v", err)
		return nil, err
	}

	res := &orders.CreateOrderResponse{Status: "Success"}
	return res, nil
}
