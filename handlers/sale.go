package handlers

import (
	"flow-app/db"
	"flow-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSale(c *gin.Context) {
	var sale models.Sale

	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the product exists and has enough quantity
	var currentQty int
	err := db.DB.QueryRow("SELECT quantity FROM products WHERE id = $1", sale.ProductID).Scan(&currentQty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}
	if sale.Quantity > currentQty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not enough inventory"})
		return
	}

	// Insert sale
	query := `INSERT INTO sales (product_id, quantity) VALUES ($1, $2) RETURNING id`
	err = db.DB.QueryRow(query, sale.ProductID, sale.Quantity).Scan(&sale.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create sale"})
		return
	}

	// Update product inventory
	_, err = db.DB.Exec("UPDATE products SET quantity = quantity - $1 WHERE id = $2", sale.Quantity, sale.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product quantity"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sale recorded", "sale": sale})
}
