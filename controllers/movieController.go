package controllers

import (
	"fmt"
	"net/http"

	"github.com/dannielss/goflix/model"
	"github.com/dannielss/goflix/repositories"
	"github.com/gin-gonic/gin"
)

func NewMovieController(movieRepository repositories.MovieRepositoryInterface) MovieControllerInterface {
	return &movieController{movieRepository}
}

type movieController struct {
	movieRepository repositories.MovieRepositoryInterface
}

type MovieControllerInterface interface {
	ShowMovies(c *gin.Context)
	AddMovie(c *gin.Context)
}

func (mc *movieController) ShowMovies(c *gin.Context) {
	movies := []model.MovieWithCategory{}

	rows, err := mc.movieRepository.ShowAll()

	if err != nil {
		fmt.Printf("Error %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var movie model.MovieWithCategory

		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Category)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})
}

func (mc *movieController) AddMovie(c *gin.Context) {
	var body model.PayloadMovie

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mc.movieRepository.Insert(&body)

	if err != nil {
		fmt.Printf("Error %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie added successfuly",
	})
}
