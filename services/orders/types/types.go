package types

import (
	"context"

	"github.com/nitheshkv96/kitchen-microservice/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
