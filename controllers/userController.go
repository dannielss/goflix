package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dannielss/goflix/model"
	"github.com/dannielss/goflix/repositories"
	"github.com/gin-gonic/gin"
)

func NewUserController(userRepository repositories.UserRepositoryInterface) UserControllerInterface {
	return &userController{userRepository}
}

type userController struct {
	userRepository repositories.UserRepositoryInterface
}

type UserControllerInterface interface {
	ShowUsers(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func (uc *userController) ShowUsers(c *gin.Context) {
	users := []model.User{}

	rows, err := uc.userRepository.ShowAll()

	if err != nil {
		fmt.Printf("Error %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong",
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Something wrong",
			})
			return
		}

		users = append(users, user)

	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uc *userController) AddUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.HashPassword()

	err := uc.userRepository.Insert(&user)

	if err != nil {
		fmt.Printf("Error %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User added successfuly",
	})
}

func (uc *userController) UpdateUser(c *gin.Context) {
	var user model.User

	idAsString := c.Param("id")

	id, err := strconv.Atoi(idAsString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Id = id

	val, error := uc.userRepository.Update(&user)

	if error != nil {
		fmt.Printf("Error %s", error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong",
		})
		return
	}

	if val == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfuly",
	})
}

func (uc *userController) DeleteUser(c *gin.Context) {
	idAsString := c.Param("id")

	id, err := strconv.Atoi(idAsString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong",
		})
		return
	}

	val, error := uc.userRepository.Delete(id)

	if error != nil {
		fmt.Printf("Error %s", error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong",
		})
		return
	}

	if val == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfuly",
	})
}
