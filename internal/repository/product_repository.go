package repository

import "github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/model"

type ProductRepository interface{
	GetAllProducts() ([]model.Product,error)
	GetProductByID(id int64) (*model.Product, error) 
	Create(product model.Product) (*model.Product, error)
	Update(id int64, product model.Product) (*model.Product, error)
	Delete(id int64) error
}