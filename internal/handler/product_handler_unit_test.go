package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	return nil, nil
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