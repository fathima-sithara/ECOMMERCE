package controllers

import (
	"net/http"
	"strconv"

	"github.com/fathima-sithara/ecommerce/auth"
	"github.com/fathima-sithara/ecommerce/database"
	"github.com/fathima-sithara/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// SignUp handles user registration
func SignUp(c *gin.Context) {
	var user models.User

	// Bind the incoming JSON to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validate user input
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := database.InitDB().Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "false",
			"Eroor":  err.Error(),
		})
	}

	// Optionally, you may generate an OTP here and send it to the user's email
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful, please proceed to OTP validation.",
	})
}

type UserLogin struct {
	Email        string
	Password     string
	Block_status bool
}

func LoginUser(c *gin.Context) {
	var userLogin UserLogin
	var user models.User

	// Bind the incoming JSON to the userLogin struct
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Fetch the user from the database by email
	if err := database.InitDB().Where("email = ?", userLogin.Email).First(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error.Error()})
		return
	}

	if !user.Verified {
		database.InitDB().Delete(&user)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":    "user is not verified. data deleted",
			"meassage": "please complete OTP verificatioon to complete registration",
		})
		return
	}

	if user.Block_Status {
		c.JSON(http.StatusForbidden, gin.H{"msg": "user has been blocked by admin"})
		return
	}

	if credentailCheckErr := user.CheckPassword(userLogin.Password); credentailCheckErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	str := strconv.Itoa(int(user.ID))
	tokenString, err := auth.GenerateJWT(str)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("UserAuth", tokenString["accessString"], 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{
		"email":    userLogin.Email,
		"password": userLogin.Password,
		"token":    tokenString["access_token"]})
}

func UserHome(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "wellcomme user home"})
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("UsereAuth", "", -1, "", "", false, false)
	c.JSON(200, gin.H{
		"message": "user sussessfully Log out",
	})
}
