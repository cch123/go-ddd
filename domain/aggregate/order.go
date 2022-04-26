package aggregate

import (
	"context"
	"github.com/cch123/go-ddd/domain/iface"
)

type OrderAgg struct {
	order iface.OrderRepo
}

func NewOrderAgg(o iface.OrderRepo) *OrderAgg {
	return &OrderAgg{order: o}
}

func (o *OrderAgg) RemoveIllegalOrder(ctx context.Context, req *RemoveIllegalOrderReq) (*RemoveIllegalOrderResponse, error) {
	// step 1, delete this order
	err := o.order.DeleteOrderByID(req.OrderID)

	// step 2, notify the user about this special event
	return nil, err
}
