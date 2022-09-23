package main

import (
	"github.com/dannielss/goflix/controller"
	"github.com/dannielss/goflix/database"
	"github.com/dannielss/goflix/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn := database.NewMySQLClient()

	defer conn.Close()

	repo := repository.NewUserRepository(conn)
	controller := controller.NewUserController(repo)

	r := gin.Default()

	r.GET("/", controller.ShowUsers)
	r.POST("/user", controller.AddUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)

	r.Run(":3333")
}
