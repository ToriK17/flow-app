package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgres://localhost:5432/flowhub?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	log.Println("Connected to PostgreSQL!")
}
