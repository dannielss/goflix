package middlewares

import (
	"github.com/dannielss/goflix/config"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(bearerSchema):]

		if !config.NewJWTConfig().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
