package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DBCon *sql.DB

func DBConnect() {
	var err error

	DBCon, err = sql.Open("mysql", "root:password@tcp(localhost)/goflix")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
}
