package models

import "time"

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
type Stock struct {
	Stock int `json:"stock" db:"stock" binding:"required"`
}
