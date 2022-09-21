package controller

import (
	"fmt"
	"net/http"

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
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User added successfuly",
	})
}
