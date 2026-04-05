package service

import (
	"context"

	"github.com/nitheshkv96/kitchen-microservice/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// store
}

// This will be injected into the handler
func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context,
	order *orders.Order) error {

	ordersDb = append(ordersDb, order)
	return nil
}
