package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dannielss/goflix/model"
	"github.com/dannielss/goflix/repository"
	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	var users []model.User

	rows, err := repository.ShowAll()

	if err != nil {
		fmt.Printf("Error %s", err)
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

func AddUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.Insert(user)

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

func UpdateUser(c *gin.Context) {
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

	val, error := repository.Update(user)

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

func DeleteUser(c *gin.Context) {
	idAsString := c.Param("id")

	id, err := strconv.Atoi(idAsString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong",
		})
		return
	}

	val, error := repository.Delete(id)

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
