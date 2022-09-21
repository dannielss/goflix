package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/goflix")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
	return db
}
