package controllers

import (
	"database/sql"
	"net/http"

	"github.com/dannielss/goflix/config"
	"github.com/dannielss/goflix/model"
	"github.com/dannielss/goflix/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewLoginController(userRepository repositories.UserRepositoryInterface) LoginControllerInterface {
	return &loginController{userRepository}
}

type loginController struct {
	userRepository repositories.UserRepositoryInterface
}

type LoginControllerInterface interface {
	Login(c *gin.Context)
}

func (lc *loginController) Login(c *gin.Context) {
	var payload model.Login

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	row := lc.userRepository.ShowOne("email = ?", payload.Email)

	scanError := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if scanError == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find user",
		})
		return
	}

	if scanError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something wrong",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	token, err := config.NewJWTConfig().GenerateToken(user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
