package order

import (
	"context"
	"github.com/cch123/go-ddd/domain/entity"
	"github.com/cch123/go-ddd/domain/irepo"
	"github.com/cch123/go-ddd/repo/data/ent"
	"github.com/cch123/go-ddd/repo/data/ent/order"
	"time"
)

// Order should impl order repo interface
type Order struct {
	client *ent.Client
} // impl irepo.OrderRepo

var _ irepo.OrderRepo = &Order{}

// NewOrder create a new customer
func NewOrder(client *ent.Client) irepo.OrderRepo {
	return &Order{client: client}
}

func (o *Order) DeleteOrderByID(ctx context.Context, orderID int) error {
	// soft delete
	err := o.client.Order.DeleteOneID(orderID).Exec(ctx)
	_, err = o.client.Order.Update().SetStatus(0).Where(order.ID(orderID)).Save(ctx)

	if err != nil {
		// wrap error
	}

	return err
}

func (o *Order) GetOrderHistoryByCustomerID(customerID int64) []entity.Order {
	return nil
}

func (o *Order) PlaceOrder(ctx context.Context, info entity.OrderCreateInfo) error {
	// query address info
	// FIXME

	// wrap logic in a transaction
	tx, err := o.client.Debug().Tx(ctx)
	if err != nil {
		return err
	}

	tx.Order.Create().
		SetStatus(1).SetAddr("Beijing").SetCustomerID(info.UserID).
		SetPrice(info.Price).SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).SaveX(ctx)
	if err != nil {
		err = tx.Rollback()
	}

	tx.Customer.Create().
		SetAddr("Heaven").SetName("alex").SetAge(31).
		SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).SaveX(ctx)
	err = tx.Commit()
	return err
}
