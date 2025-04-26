package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/controllers"
	"github.com/fathimasithara01/ecommerce/middleware"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/signup", controllers.SignUp)
		user.POST("/signup/otp", controllers.ValidateOTPHandler)

		user.POST("/login", controllers.LoginUser)
		user.GET("/logout", middleware.UserAuth(), controllers.LogoutUser)
	}
}
