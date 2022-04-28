package aggregate

import (
	"context"
	"fmt"
	"github.com/cch123/go-ddd/domain/entity"

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
	err := o.order.DeleteOrderByID(ctx, req.OrderID)

	// step 2, notify the user about this special event
	fmt.Println(o.order.PlaceOrder(ctx,
		entity.OrderCreateInfo{UserID: 1, AddrID: 20,
			SkuList: []entity.SkuItem{{SkuID: "potata_level_a", Count: 10}}, Price: 123.05}))
	return nil, err
}
