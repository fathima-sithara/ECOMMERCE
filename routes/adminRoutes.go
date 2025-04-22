package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/fathima-sithara/ecommerce/controllers"
	"github.com/fathima-sithara/ecommerce/middleware"
)

func AdminRoutes(c *gin.Engine) {
	admin := c.Group("/admin")
	{
		admin.POST("/signup", controllers.AdminSignup)
		admin.POST("/login", controllers.AdminLogin)
		admin.GET("/home", middleware.AdminAuth(), controllers.AdminHome)
	}
}
