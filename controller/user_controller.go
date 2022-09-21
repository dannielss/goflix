package controller

import (
	"fmt"
	"net/http"

	"github.com/dannielss/goflix/database"
	"github.com/dannielss/goflix/model"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	var users []model.User

	rows, err := database.DBCon.Query("SELECT * FROM users")

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
