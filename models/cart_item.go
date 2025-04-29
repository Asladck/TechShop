package models

import "time"

type CartItem struct {
	Id        string    `json:"cart_id" db:"cart_id"`
	UserID    string    `json:"user_id" db:"user_id"`
	ItemID    string    `json:"item_id" binding:"required" db:"cart_item_id"`
	ItemCount int       `json:"item_count" binding:"required" db:"item_count"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"cart_created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"cart_updated_at"`
}
type Cart struct {
	CartItem
	Item
}
type CartUpdate struct {
	ItemCount int `json:"item_count" binding:"required" db:"item_count"`
}
