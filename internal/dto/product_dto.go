package dto

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Price       int64   `json:"price" binding:"required,gte=0"`
	Stock       int32   `json:"stock" binding:"required,gte=0"`
	Category    string  `json:"category" binding:"required"`
	IsActive    *bool   `json:"is_active"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Price       int64   `json:"price" binding:"required,gte=0"`
	Stock       int32   `json:"stock" binding:"required,gte=0"`
	Category    string  `json:"category" binding:"required"`
	IsActive    *bool   `json:"is_active"`
}