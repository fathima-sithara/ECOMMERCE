package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/src/controllers"
	// "github.com/fathimasithara01/ecommerce/utils/middleware"
)

func UserRoutes(v1 *gin.RouterGroup) {

	v1.POST("/register", controllers.RegisterUser)
	v1.POST("/login", controllers.UserAuthLogin)
	// v1.POST("/verify-otp",controllers.OtpVerifiecation)
	v1.POST("/sent-otp", controllers.SentOtp)
	v1.POST("/send-otp", controllers.ForgotPassword)
	v1.POST("/reset-password", controllers.UserNameValidation)
	// v1.GET("/logout", middleware.UserAuth()(), controllers.LogoutUser)

}
