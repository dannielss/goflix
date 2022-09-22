package main

import (
	"github.com/dannielss/goflix/controller"
	"github.com/dannielss/goflix/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.DBConnect()

	defer database.DBCon.Close()

	r := gin.Default()

	r.GET("/", controller.ShowUsers)
	r.POST("/user", controller.AddUser)
	r.DELETE("/user/:id", controller.DeleteUser)

	r.Run(":3333")
}
