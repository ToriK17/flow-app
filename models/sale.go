package models

type Sale struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type MonthlySales struct {
	Category   string  `json:"category"`
	TotalSales float64 `json:"total_sales"`
}
