package aggregate

import "github.com/cch123/go-ddd/domain/iface"

type CustomerAgg struct {
	customer iface.CustomerRepo
}

func NewCustomerAgg(customer iface.CustomerRepo) *CustomerAgg {
	return &CustomerAgg{
		customer: customer,
	}
}
