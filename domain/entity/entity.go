package entity

import "time"

// Order is an entity
type Order struct {
	ID        int64
	UserID    int64
	CreateAt  time.Time
	UpdatedAt time.Time

	Details []OrderDetail
}

// OrderDetail is a value object
type OrderDetail struct {
	SKU string
}
