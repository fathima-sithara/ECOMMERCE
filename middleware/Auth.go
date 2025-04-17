package middleware

import (
	"github.com/fathima-sithara/ecommerce/auth"
	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("UserAuth")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
		}

		err = auth.VlidateToken(tokenString)
		c.Set("userid", auth.P)

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
