package models

import "time"

// Wishlist represents user's wishlist item
// @Description Product saved in user's wishlist
type Wishlist struct {
	Id        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	ItemID    string    `json:"item_id" db:"item_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
