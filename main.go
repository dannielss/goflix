package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbConnect()

	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.Run(":3333")
}

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/goflix")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
	return db
}
