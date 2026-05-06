package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       int64   `json:"price" validate:"required,gte=0"`
	Stock       int32   `json:"stock" validate:"required,gte=0"`
	Category    string  `json:"category" validate:"required"`
	IsActive    *bool   `json:"is_active"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       int64   `json:"price" validate:"required,gte=0"`
	Stock       int32   `json:"stock" validate:"required,gte=0"`
	Category    string  `json:"category" validate:"required"`
	IsActive    *bool   `json:"is_active"`
}