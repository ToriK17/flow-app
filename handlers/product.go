package handlers

import (
	"flow-app/db"
	"flow-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// c *gin.Context is like a req object but more so, like a request-response context
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO products (name, category, price, quantity) 
              VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.DB.QueryRow(query, product.Name, product.Category, product.Price, product.Quantity).Scan(&product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": product})
}

func ListProducts(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, category, price, quantity FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan product"})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
