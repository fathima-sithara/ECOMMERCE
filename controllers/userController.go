package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/fathima-sithara/ecommerce/auth"
	"github.com/fathima-sithara/ecommerce/database"
	"github.com/fathima-sithara/ecommerce/models"
	"github.com/fathima-sithara/ecommerce/utils"
)

var validate = validator.New()

// SignUp handles user registration
func SignUp(c *gin.Context) {
	var user models.User

	// Bind incoming JSON to struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Validate input using struct tags
	if err := validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	// Hash user password
	if err := utils.UserHashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	otp, err := utils.VerifyAndSendOTP(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP", "details": err.Error()})
		return
	}
	user.Otp = otp
	// Save user to database
	if err := database.InitDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Regitration successful. Please verify your account via OTP.",
	})
}

// UserLogin is the login input payload
type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginUser handles user login and JWT generation
func LoginUser(c *gin.Context) {
	var input UserLogin
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Validate input
	if err := validate.Struct(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error", "details": err.Error()})
		return
	}

	db := database.InitDB()
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Check verification and block status
	if !user.Verified {
		db.Delete(&user)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "User not verified. Account deleted.",
			"message": "Please complete OTP verification.",
		})
		return
	}

	if user.Block_Status {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is blocked by admin"})
		return
	}

	// Check password
	if err := utils.UserCheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	userID := strconv.Itoa(int(user.ID))
	tokens, err := auth.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set secure cookie
	c.SetCookie("UserAuth", tokens["access_token"], 3600*24*30, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokens["access_token"],
	})
}

// UserHome serves the protected home endpoint
func UserHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to user home"})
}

// LogoutUser clears the auth cookie
func LogoutUser(c *gin.Context) {
	// Clear cookie
	c.SetCookie("UserAuth", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
