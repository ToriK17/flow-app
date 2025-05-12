package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Product struct {
    Name     string  `json:"name"`
    Category string  `json:"category"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}

func CreateProduct(c *gin.Context) {
    var product Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // DB insert logic here
    c.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": product})
}
