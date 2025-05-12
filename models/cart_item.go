package models

import "time"

// CartItem represents an item in user's shopping cart
// @Description Shopping cart item information
type CartItem struct {
	Id        string    `json:"cart_id" db:"cart_id"`
	UserID    string    `json:"user_id" db:"user_id"`
	ItemID    string    `json:"item_id" binding:"required" db:"cart_item_id"`
	ItemCount int       `json:"item_count" binding:"required" db:"item_count"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"cart_created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"cart_updated_at"`
}

// Cart represents full cart information with item details
// @Description Full shopping cart information including product details
type Cart struct {
	CartItem
	Item
}

// CartUpdate represents data for updating cart item quantity
// @Description Data structure for updating item quantity in cart
type CartUpdate struct {
	ItemCount int `json:"item_count" binding:"required" db:"item_count"`
}
