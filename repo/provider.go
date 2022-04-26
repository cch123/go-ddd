package repo

import (
	"github.com/cch123/go-ddd/repo/customer"
	"github.com/cch123/go-ddd/repo/order"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(customer.NewCustomer, order.NewOrder)
