package controller

import (
	"fmt"
	"net/http"

	"github.com/dannielss/goflix/model"
	"github.com/dannielss/goflix/repository"
	"github.com/gin-gonic/gin"
)

func NewCategoryController(categoryRepository repository.CategoryRepositoryInterface) CategoryControllerInterface {
	return &categoryController{categoryRepository}
}

type categoryController struct {
	categoryRepository repository.CategoryRepositoryInterface
}

type CategoryControllerInterface interface {
	ShowCategories(c *gin.Context)
	AddCategory(c *gin.Context)
}

func (cc *categoryController) ShowCategories(c *gin.Context) {
	categories := []model.Category{}

	rows, err := cc.categoryRepository.ShowAll()

	if err != nil {
		fmt.Printf("Error %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var category model.Category

		err := rows.Scan(&category.Id, &category.Name)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (cc *categoryController) AddCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.categoryRepository.AddCategory(&category)

	if err != nil {
		fmt.Printf("Error %s", err.Error())
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category added successfuly",
	})
}
