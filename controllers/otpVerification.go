package controllers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/fathima-sithara/ecommerce/database"
	"github.com/fathima-sithara/ecommerce/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// / generateOTP creates a random 4-digit OTP
func generateOTP() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		return "", err
	}
	// sendEmail(email)
	return strconv.FormatInt(n.Int64()+1000, 10), nil
}

// sendEmail sends an OTP to the provided email address
func sendEmail(toEmail, otp string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	body := fmt.Sprintf("Subject: OTP Verification\n\nYour OTP is: %s", otp)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, []byte(body))
}

// VerifyAndSendOTP generates and sends an OTP to the user's email
func VerifyAndSendOTP(email string) (string, error) {
	otp, err := generateOTP()
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}

	if err := sendEmail(email, otp); err != nil {
		return "", fmt.Errorf("failed to send email: %w", err)
	}

	return otp, nil
}

// ValidateOTPHandler verifies the OTP entered by the user
func ValidateOTPHandler(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := database.InitDB()
	var user models.User
	if err := db.Where("email = ? AND otp = ?", input.Email, input.OTP).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or OTP"})
		return
	}

	user.Verified = true
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update verification status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully. Account activated."})
}

// SendForgotPasswordOTPHandler sends OTP to user email for password reset
func SendForgotPasswordOTPHandler(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db := database.InitDB()
	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	otp, err := VerifyAndSendOTP(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	if err := db.Model(&user).Update("otp", otp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update OTP in database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully. Proceed to change password."})
}

// ChangePasswordHandler allows the user to change password using OTP
func ChangePasswordHandler(c *gin.Context) {
	var input struct {
		Email           string `json:"email"`
		OTP             string `json:"otp"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	db := database.InitDB()
	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if input.OTP != user.Otp {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if err := db.Model(&user).Update("password", hashedPassword).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
