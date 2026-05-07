package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"errors"

	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//mock repository
type MockProductRepository struct{}

func (m *MockProductRepository) GetAllProducts() ([]model.Product, error) {
	return []model.Product{
		{
			ID:          1,
			Name:        "Indomie",
			Description: stringPtr("Instant noodle"),
			Price:       3500,
			Stock:       100,
			Category:    "Food",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (m *MockProductRepository) GetProductByID(id int64) (*model.Product, error) {
	if id == 1 {
		return &model.Product{
			ID:          1,
			Name:        "Indomie",
			Description: stringPtr("Instant noodle"),
			Price:       3500,
			Stock:       100,
			Category:    "Food",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, nil
	}

	return nil, errors.New("product not found")
}

func (m *MockProductRepository) Create(product model.Product) (*model.Product, error) {
	return nil, nil
}

func (m *MockProductRepository) Update(id int64, product model.Product) (*model.Product, error) {
	return nil, nil
}

func (m *MockProductRepository) Delete(id int64) error {
	return nil
}


func stringPtr(s string) *string {
	return &s
}

// unit test untuk endpoint GET /products
func TestGetAllProducts(t *testing.T) {

	gin.SetMode(gin.TestMode)
	mockRepo := &MockProductRepository{}

	productHandler := NewProductHandler(mockRepo)

	router := gin.Default()
	router.GET("/products", productHandler.GetAllProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string][]model.Product

	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Len(t, response["data"], 1)
	assert.Equal(t, "Indomie", response["data"][0].Name)
}


//unit test untuk endpoint GET /products/:id (fetch produk yang idnya ada)

func TestGetProductByID_success(t *testing.T){
	gin.SetMode(gin.TestMode)

	mockRepo := &MockProductRepository{}
	productHandler := NewProductHandler(mockRepo)

	router := gin.Default()
	router.GET("/products/:id", productHandler.GetProductByID)

	//cek untuk product yang idnya ada
	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]model.Product

	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(t, int64(1), response["data"].ID)
	assert.Equal(t, "Indomie", response["data"].Name)
}

//unit test untuk endpoint GET /products/:id (fetch produk yang idnya valid tapi tidak ada)
func TestGetProductByID_NotFound(t *testing.T){
	gin.SetMode(gin.TestMode)

	mockRepo := &MockProductRepository{}
	productHandler := NewProductHandler(mockRepo)

	router := gin.Default()
	router.GET("/products/:id", productHandler.GetProductByID)
	//cek untuk product yang idnya tidak ada

	req, _ := http.NewRequest(http.MethodGet, "/products/67", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)

	var errorResponse map[string]string

	err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, "product not found", errorResponse["error"])
}

//unit test untuk endpoint GET /products/:id (fetch produk yang idnya tidak valid)
func TestGetProductByID_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &MockProductRepository{}
	productHandler := NewProductHandler(mockRepo)

	router := gin.Default()
	router.GET("/products/:id", productHandler.GetProductByID)

	req, _ := http.NewRequest(http.MethodGet, "/products/abc", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	var errorResponse map[string]string

	err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, "invalid id", errorResponse["error"])
}