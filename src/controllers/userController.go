package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/src/services"
	"github.com/fathimasithara01/ecommerce/utils/constant"
	"github.com/fathimasithara01/ecommerce/utils/response"
	validator "github.com/fathimasithara01/ecommerce/utils/validation"
)

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"required" `
}

// Signup
func RegisterUser(c *gin.Context) {
	var request RegisterUserRequest

	// Bind JSON and validate fields
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	err := validator.Validate(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	// Call user registration service
	userService := services.UserServices{}
	token, err := userService.Register(request.Username, request.Email, request.Password, request.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}

	// Success response

	c.JSON(http.StatusOK, response.SuccessResponse(map[string]interface{}{
		"message":  "User registered successfully",
		"message2": "pleas veify your email",
		"token":    token,
	}))

}

// User Login
func UserAuthLogin(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	err := validator.Validate(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	userService := services.UserServices{}
	localToken, u, formsubmitted, err := userService.Login(request.Email, request.Password)
	if err != nil {
		// Check if the error is related to a missing user (custom error message)
		if err.Error() == "We couldn't find any account associated with this email." {
			c.JSON(http.StatusUnauthorized, response.ErrorMessage(constant.UNAUTHORIZED, err))
			return
		}

	}
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]interface{}{
		"local_token":    localToken,
		"form_submitted": formsubmitted,
		"user":           u,
	}))
}
func SentOtp(c *gin.Context) {
	meth := c.Query("meth")
	if meth == "" {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("pleas enter methord")))
		return
	}

	var request struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	err := validator.Validate(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	userService := services.UserServices{}
	token, res, err := userService.OtpService(request.Email, meth)
	if err != nil {
		// Check if the error is related to a missing user (custom error message)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]interface{}{
		"message": res,
		"token":   token,
	}))
}
func ForgotPassword(c *gin.Context) {
	token := c.Query("token")
	var request struct {
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {

		// Return any other validation errors
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	err := validator.Validate(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	userService := services.UserServices{}
	err = userService.ForgotPassword(request.Password, token)
	if err != nil {
		// Check if the error is related to a missing user (custom error message)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(map[string]interface{}{
		"message": "Password reset successfully",
	}))
}

func UserNameValidation(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	err := validator.Validate(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	userService := services.UserServices{}
	err = userService.ValidateUserName(request.Username)
	if err != nil {
		// Check if the error is related to a missing user (custom error message)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(map[string]interface{}{
		"message": "you can use this user name",
	}))
}

// UserHome serves the protected home endpoint
func UserHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to user home"})
}

// LogoutUser clears the utils cookie
func LogoutUser(c *gin.Context) {
	// Clear cookie
	c.SetCookie("Userutils", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
