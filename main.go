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

	movieRepo := repository.NewMovieRepository(conn)
	movieController := controller.NewMovieController(movieRepo)

	r := gin.Default()

	r.GET("/", userController.ShowUsers)
	r.POST("/user", userController.AddUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	r.GET("/categories", categoryController.ShowCategories)
	r.POST("/category", categoryController.AddCategory)

	r.GET("/movies", movieController.ShowMovies)
	r.POST("/movie", movieController.AddMovie)

	r.Run(":3333")
}
