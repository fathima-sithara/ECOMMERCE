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

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	validation := validate.Struct(user)
	if validation != nil {
		c.JSON(404, gin.H{"error": validation})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// otp := VerifyOTP(user.Email)
	// if result := database.InitDB().Create(&user).Error; result != nil {
	// 	c.JSON(500, gin.H{
	// 		"status": "false",
	// 		"Eroor":  result.Error(),
	// 	})
	// } else {
	// 	database.InitDB().Model(&user).Where("email LIKE ?", user.Email).Update("otp", otp)
	// 	c.JSON(200, gin.H{
	// 		"message": "Go to otp validation",
	// 	})
	// }

}

type UserLogin struct {
	Email        string
	Password     string
	Block_status bool
}

func LoginUser(c *gin.Context) {
	var userLogin UserLogin
	var user models.User

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.InitDB().Raw("select *from users where email=?", userLogin.Email).Scan(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	if !user.Verified {
		db := database.InitDB()
		db.Delete(&user)

		c.JSON(422, gin.H{
			"error":    "user is not verified. data deleted",
			"meassage": "please complete OTP verificatioon to complete registration",
		})
		return
	}

	if user.Block_Status {
		c.JSON(404, gin.H{"msg": "user has been blocked by admin"})
		return
	}

	credentailCheck := user.CheckPassword(userLogin.Password)
	if credentailCheck != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.Abort()
		return
	}

	str := strconv.Itoa(int(user.ID))
	tokenString, err := auth.GenerateJWT(str)
	token := tokenString["acess_token"]
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", token, 3600*24*30, "", "", false, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"email": userLogin.Email, "password": userLogin.Password, "token": tokenString})
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
