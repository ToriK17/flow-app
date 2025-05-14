# 🌿 Flow-App – POS Inventory Microservice

### 🚀 What This Is

**Flow-app** is a backend API built in **Go** using **Gin** and **PostgreSQL**, designed to simulate a compliant inventory and point-of-sale system for a standardized business.

This project demonstrates:

- Translating business logic into functional backend systems
- Handling core retail and compliance-style workflows (inventory, sales, reporting)


### 🌟 Features

- **Add Products** `POST /products`  
Create inventory items with name, category, price, and quantity.  

- **List Products** `GET /products`
Fetch all inventory items.

- **Record Sales** `POST /sales`
Simulates a real POS transaction and automatically deducts inventory.

- **Monthly Sales Metrics** `GET /metrics/monthly-sales`
Returns revenue totals by category for the current month.

- **Unsellable Products** `GET /metrics/unsellables`
Flags anything that's been sitting unsold for 60+ days.

### 🛠️ Tech Stack
- Go
- Gin – Fast HTTP router and middleware
- PostgreSQL
- golang-migrate – Schema migrations like a real adult (yes, I miss Rails too)
- curl/Postman – To hit those endpoints

### 🧪 Running It
Start the server: `go run main.go`
Migrate the database: `migrate -path db/migrations -database "postgres://localhost:5432/flowhub?sslmode=disable" up`

### ✨ Why This Exists
Because I wanted to see what I could do in a few nights after work to expose myself to GO and it's tooling. This isn’t a full system, but it reflects my ability to prioritize, learn, and create functional backend architecture for funsies.

### 📜 License
MIT. Free to use, but you still gotta be cool about it.
