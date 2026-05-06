package repository

import (
	"database/sql"
	"errors"

	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}


func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product

	query := `
		SELECT id, name, description, price, stock, category, is_active, created_at, updated_at
		FROM products
		ORDER BY id DESC
	`

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// =========================
func (r *productRepository) GetProductByID(id int64) (*model.Product, error) {
	var product model.Product

	query := `
		SELECT id, name, description, price, stock, category, is_active, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	err := r.db.Get(&product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Create(product model.Product) (*model.Product, error) {
	query := `
		INSERT INTO products (name, description, price, stock, category, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, description, price, stock, category, is_active, created_at, updated_at
	`

	var created model.Product

	err := r.db.Get(
		&created,
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.Category,
		product.IsActive,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *productRepository) Update(id int64, product model.Product) (*model.Product, error) {
	query := `
		UPDATE products
		SET name = $1,
			description = $2,
			price = $3,
			stock = $4,
			category = $5,
			is_active = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING id, name, description, price, stock, category, is_active, created_at, updated_at
	`

	var updated model.Product

	err := r.db.Get(
		&updated,
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.Category,
		product.IsActive,
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &updated, nil
}

func (r *productRepository) Delete(id int64) error {
	query := `DELETE FROM products WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}