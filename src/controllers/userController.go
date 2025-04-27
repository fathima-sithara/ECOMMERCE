package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/src/models"
	"github.com/fathimasithara01/ecommerce/src/services"
	"github.com/fathimasithara01/ecommerce/utils"
	"github.com/fathimasithara01/ecommerce/utils/constant"
	"github.com/fathimasithara01/ecommerce/utils/jwt"
	"github.com/fathimasithara01/ecommerce/utils/response"
	validator "github.com/fathimasithara01/ecommerce/utils/validation"
)

type User struct {
	FirstName string `json:"first_name" gorm:"not null" `
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique" validate:"required"`
	Password  string `json:"password" gorm:"not null" validate:"required"`
	Phone     string `json:"phone" gorm:"not null;" `
}

func SignUp(c *gin.Context) {
	var req User

	// Bind incoming JSON to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.BADREQUEST, err))

		return
	}

	err := validator.Validate(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.BADREQUEST, err))
	}

	us := services.UserServices{}
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
	}
	if err := us.RegisterUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(constant.INTERNALSERVERERROR, err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponseMsg(map[string]interface{}{
		"data": "",
	}, "succes fully registration"))
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func LoginUser(c *gin.Context) {
	var input UserLogin
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// if err := validate.Struct(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error", "details": err.Error()})
	// 	return
	// }

	db := database.PgSQLDB
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !user.Verified {
		db.Delete(&user)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "User not verified. Account deleted.",
			"message": "Please complete OTP verification.",
		})
		return
	}

	// if user.Block_Status {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "User is blocked by admin"})
	// 	return
	// }

	// Check password
	if err := utils.UserCheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	userID := strconv.Itoa(int(user.ID))
	tokens, err := jwt.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set secure cookie
	c.SetCookie("Userutils", tokens["access_token"], 3600*24*30, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokens["access_token"],
	})
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
