package routes

import (
	"github.com/fathimasithara01/ecommerce/src/controllers"
	"github.com/fathimasithara01/ecommerce/utils/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("/signup", controllers.SignUp)
		user.POST("/signup/otp", controllers.ValidateOTPHandler)

		user.POST("/login", controllers.LoginUser)
		user.GET("/logout", middleware.UserAuth(), controllers.LogoutUser)
	}
}
