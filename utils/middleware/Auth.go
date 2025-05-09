package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/fathimasithara01/ecommerce/utils/jwt"
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
		// err = jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// userId := jwt.P
		// c.Set("userid", userId)
		c.Next()

	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStirng, err := c.Cookie("AdminAuth")
		if tokenStirng == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not ccontain  an access token"})
			c.Abort()
			return
		}
		// err = jwt.ValidateToken(tokenStirng)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
