package customer

import (
	"github.com/cch123/go-ddd/domain/irepo"
)

// Customer should impl customer repo interface
type Customer struct{} // impl Customer repo
var _ irepo.CustomerRepo = &Customer{}

// NewCustomer create a new customer
func NewCustomer() irepo.CustomerRepo {
	return &Customer{}
}

func (c *Customer) UpdateEmail(email string) error {
	return nil
}

func (c *Customer) RelocateTo(addr string) error {
	return nil
}
