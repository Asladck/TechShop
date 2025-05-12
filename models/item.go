package models

import "time"

// Item represents a product in the store
// @Description Product information
type Item struct {
	Id          string    `json:"id" db:"id"`
	Title       string    `json:"title" binding:"required" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Price       float64   `json:"price" binding:"required" db:"price"`
	Stock       int       `json:"stock" db:"stock"`
	ImageURL    string    `json:"image_url,omitempty" db:"image_url"`
	CreatedAt   time.Time `json:"item_created_at,omitempty" db:"item_created_at"`
	UpdatedAt   time.Time `json:"item_updated_at,omitempty" db:"item_updated_at"`
}

// Stock represents inventory stock level
// @Description Product stock information for updates
type Stock struct {
	Stock int `json:"stock" db:"stock" binding:"required"`
}
