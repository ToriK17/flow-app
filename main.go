package main

import (
	"flowhub/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.POST("/products", CreateProduct)
	r.GET("/products", ListProducts)

	r.POST("/sales", CreateSale)
	r.GET("/metrics/monthly-sales", GetMonthlySales)
	r.GET("/metrics/unsellables", GetUnsellableProducts)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
