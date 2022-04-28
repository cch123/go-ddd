package entity

import "time"

// Order is an entity
type Order struct {
	ID        int
	UserID    int64
	CreateAt  time.Time
	UpdatedAt time.Time

	Details []OrderDetail
}

// OrderDetail is a value object
type OrderDetail struct {
	SkuID int
	Count int
}

type OrderCreateInfo struct {
	UserID  int
	AddrID  int
	SkuList []SkuItem
	Price   float64
}

type SkuItem struct {
	SkuID string
	Count int
}
