package model

import "time"

type Product struct {
	ID int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Price int64 `json:"price" db:"price"`
	Stock int32 `json:"stock" db:"stock"`
	Category string `json:"category" db:"category"`
	IsActive bool `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"` 
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` 
}
