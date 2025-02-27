package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Colocar em variaveis de ambiente os.Getenv("TOKEN")
const validToken = "teste123"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] != validToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
