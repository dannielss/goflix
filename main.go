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

	userRepo := repository.NewUserRepository(conn)
	userController := controller.NewUserController(userRepo)

	categoryRepo := repository.NewCategoryRepository(conn)
	categoryController := controller.NewCategoryController(categoryRepo)

	r := gin.Default()

	r.GET("/", userController.ShowUsers)
	r.POST("/user", userController.AddUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	r.GET("/categories", categoryController.ShowCategories)

	r.Run(":3333")
}
