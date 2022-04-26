package order

import (
	"fmt"
	"github.com/cch123/go-ddd/domain/entity"
	"github.com/cch123/go-ddd/domain/iface"
	"time"
)

// Order should impl order repo interface
type Order struct{} // impl iface.OrderRepo

var _ iface.OrderRepo = &Order{}

// NewOrder create a new customer
func NewOrder() iface.OrderRepo {
	return &Order{}
}

func (o *Order) DeleteOrderByID(orderID int64) error {
	fmt.Println(orderID)
	return nil
}

func (o *Order) GetTopAmountOrder() entity.Order {
	return entity.Order{
		ID:        1,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
		Details:   []entity.OrderDetail{{SKU: "Bleach"}},
	}
}
func (o *Order) GetOrderHistoryByCustomerID(customerID int64) []entity.Order {
	return nil
}
