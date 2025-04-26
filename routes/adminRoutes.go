package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/controllers"
	"github.com/fathimasithara01/ecommerce/middleware"
)

func AdminRoutes(c *gin.Engine) {
	admin := c.Group("/admin")
	{
		admin.POST("/signup", controllers.AdminSignup)
		admin.POST("/login", controllers.AdminLogin)
		admin.GET("/home", middleware.AdminAuth(), controllers.AdminHome)
	}
}
