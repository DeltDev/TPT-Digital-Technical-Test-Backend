package main

import (
	"log"

	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/db"
	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/handler"
	"github.com/DeltDev/TPT-Digital-Technical-Test-Backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	database := db.NewDB()

	productRepo := repository.NewProductRepository(database)

	productHandler := handler.NewProductHandler(productRepo)

	r := gin.Default()

	r.GET("/products", productHandler.GetAllProducts)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.POST("/products", productHandler.CreateProduct)
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}