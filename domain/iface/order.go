package iface

import "github.com/cch123/go-ddd/domain/entity"

// OrderRepo is for order iface
type OrderRepo interface {
	// GetTopAmountOrder get the most expensive order
	GetTopAmountOrder() entity.Order

	DeleteOrderByID(orderID int64) error

	GetOrderHistoryByCustomerID(customerID int64) []entity.Order
}
