package middleware

import (
	"net/http"

	"github.com/fathima-sithara/ecommerce/auth"
	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the request cookie
		tokenString, err := c.Cookie("UserAuth")
		if err != nil || tokenString == "" {
			// If no token found, return an unauthorized response
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain a valid access token"})
			c.Abort()
			return
		}
		err = auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		userId := auth.P
		c.Set("userid", userId)
		c.Next()
	}
}
