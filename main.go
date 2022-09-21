package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dannielss/goflix/controller"
	"github.com/dannielss/goflix/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnect()

	defer database.DBCon.Close()
	r := gin.Default()

	r.GET("/", controller.Show)

	r.Run(":3333")
}

func dbConnect() {
	var err error

	database.DBCon, err = sql.Open("mysql", "root:password@tcp(localhost)/goflix")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
}
