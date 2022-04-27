package irepo

// CustomerRepo is for customer irepo
type CustomerRepo interface {
	UpdateEmail(email string) error // UpdateEmail update customer's email
	RelocateTo(addr string) error
}
