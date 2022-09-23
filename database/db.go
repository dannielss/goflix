package database

import (
	"database/sql"
	"fmt"
	"log"
)

func NewMySQLClient() *sql.DB {
	conn, err := sql.Open("mysql", "root:password@tcp(localhost)/goflix")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")

	return conn
}
