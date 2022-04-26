package iface

// CustomerRepo is for customer iface
type CustomerRepo interface {
	UpdateEmail(email string) error // UpdateEmail update customer's email
	RelocateTo(addr string) error
}
