package models

import "time"

type Order struct {
	Id        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	ItemID    string    `json:"item_id" db:"item_id"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	Status    string    `json:"status" db:"status"`
	ItemCount int       `json:"item_count" db:"item_count"` // ordered, delivering, delivered, cancelled
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
type OrderItem struct {
	Id              string  `json:"id" db:"id"`
	OrderID         string  `json:"order_id" db:"order_id"`
	ItemID          string  `json:"item_id" db:"item_id"`
	ItemCount       int     `json:"item_count" db:"item_count"`
	PriceAtPurchase float64 `json:"price_at_purchase" db:"price_at_purchase"`
}
