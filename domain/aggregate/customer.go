package aggregate

import "github.com/cch123/go-ddd/domain/irepo"

type CustomerAgg struct {
	customer irepo.CustomerRepo
}

func NewCustomerAgg(customer irepo.CustomerRepo) *CustomerAgg {
	return &CustomerAgg{
		customer: customer,
	}
}
