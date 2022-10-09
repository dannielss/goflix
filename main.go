package main

import (
	"github.com/dannielss/goflix/controllers"
	"github.com/dannielss/goflix/database"
	"github.com/dannielss/goflix/middlewares"
	"github.com/dannielss/goflix/repositories"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn := database.NewMySQLClient()

	defer conn.Close()

	userRepo := repositories.NewUserRepository(conn)
	userController := controllers.NewUserController(userRepo)
	loginController := controllers.NewLoginController(userRepo)

	categoryRepo := repositories.NewCategoryRepository(conn)
	categoryController := controllers.NewCategoryController(categoryRepo)

	movieRepo := repositories.NewMovieRepository(conn)
	movieController := controllers.NewMovieController(movieRepo)

	r := gin.Default()

	r.POST("/login", loginController.Login)
	r.POST("/user", userController.AddUser)

	r.Use(middlewares.Auth())

	r.GET("/users", userController.ShowUsers)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	r.GET("/categories", categoryController.ShowCategories)
	r.POST("/category", categoryController.AddCategory)

	r.GET("/movies", movieController.ShowMovies)
	r.POST("/movie", movieController.AddMovie)

	r.Run(":3333")
}
