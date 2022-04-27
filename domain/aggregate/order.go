package aggregate

import (
	"context"
	"github.com/cch123/go-ddd/domain/irepo"
)

type OrderAgg struct {
	order irepo.OrderRepo
}

func NewOrderAgg(o irepo.OrderRepo) *OrderAgg {
	return &OrderAgg{order: o}
}

func (o *OrderAgg) RemoveIllegalOrder(ctx context.Context, req *RemoveIllegalOrderReq) (*RemoveIllegalOrderResponse, error) {
	// step 1, delete this order
	err := o.order.DeleteOrderByID(req.OrderID)

	// step 2, notify the user about this special event
	return nil, err
}
