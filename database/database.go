package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		fmt.Println("DATABASE_URL is not set")
		return
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	DB = db
	fmt.Println("Database connected successfully!")
}