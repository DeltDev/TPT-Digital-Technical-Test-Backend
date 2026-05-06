package handler

import (
	"net/http"
	"strconv"

	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/dto"
	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/model"
	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// endpoint GET /products

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.repo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

//endpoint GET /products/:id

func (h *ProductHandler) GetProductByID(c *gin.Context){
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, errFetch := h.repo.GetProductByID(int64(idInt))

	if errFetch != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errFetch.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// endpoint POST /products

func (h *ProductHandler) CreateProduct(c *gin.Context){
	var req dto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		IsActive:    isActive,
	}

	created, err := h.repo.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": created,
	})
}