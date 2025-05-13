package handlers

import (
	"flow-app/db"
	"flow-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMonthlySales(c *gin.Context) {
	query := `
		SELECT p.category, SUM(s.quantity * p.price) AS total_sales
		FROM sales s
		JOIN products p ON s.product_id = p.id
		WHERE DATE_TRUNC('month', s.sale_date) = DATE_TRUNC('month', CURRENT_DATE)
		GROUP BY p.category
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch metrics"})
		return
	}
	defer rows.Close()

	var metrics []models.MonthlySales

	for rows.Next() {
		var m models.MonthlySales
		if err := rows.Scan(&m.Category, &m.TotalSales); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan metrics"})
			return
		}
		metrics = append(metrics, m)
	}

	c.JSON(http.StatusOK, gin.H{"monthly_sales": metrics})
}

func GetUnsellableProducts(c *gin.Context) {
	query := `
		SELECT id, name, category, price, quantity
		FROM products
		WHERE created_at < NOW() - INTERVAL '60 days';
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch unsellable products"})
		return
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan product"})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, gin.H{"unsellable_products": products})
}
