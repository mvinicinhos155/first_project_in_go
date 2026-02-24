package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

	var DB *sql.DB

func Connect() {

	dns := "root:@tcp(localhost:3307)/first_project"

	db, err := sql.Open("mysql", dns)

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Error pinging database: ", err)
	}

	DB = db

	fmt.Println("Database connected successfully!")
}