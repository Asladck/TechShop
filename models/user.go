package models

import "time"

type User struct {
	Id        string    `json:"-" db:"id"`
	Email     string    `json:"email" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password_hash" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
