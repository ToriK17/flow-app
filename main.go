package main

import (
	"flow-app/db"
	"flow-app/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.POST("/products", handlers.CreateProduct)

	r.GET("/products", handlers.ListProducts)

	r.POST("/sales", handlers.CreateSale)

	r.GET("/metrics/monthly-sales", handlers.GetMonthlySales)

	r.GET("/metrics/unsellables", handlers.GetUnsellableProducts)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
