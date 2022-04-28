package irepo

import (
	"context"
	"github.com/cch123/go-ddd/domain/entity"
)

// OrderRepo is for order irepo
type OrderRepo interface {
	DeleteOrderByID(ctx context.Context, orderID int) error
	GetOrderHistoryByCustomerID(customerID int64) []entity.Order
	PlaceOrder(ctx context.Context, info entity.OrderCreateInfo) error
}
