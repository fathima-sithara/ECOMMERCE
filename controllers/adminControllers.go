package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/models"
	"github.com/fathimasithara01/ecommerce/utils"
)

type AdminLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
}

type OrderResponse struct {
	OrderID     uint      `json:"order_id"`
	UserID      uint      `json:"user_id"`
	Status      string    `json:"status"`
	PaymentMode string    `json:"payment_mode"`
	TotalAmount uint      `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

// AdminSignup handles admin registration
func AdminSignup(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.InitDB()
	if err := db.Where("email = ?", admin.Email).First(&models.Admin{}).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin already exists with this email"})
		return
	}

	if err := utils.AdminHashPassword(admin.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := db.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin created successfully"})
}

// AdminLogin utilsenticates an admin and returns JWT token
func AdminLogin(c *gin.Context) {
	var req AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.InitDB()
	var admin models.Admin
	if err := db.Where("email = ?", req.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := utils.AdminCheckPassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenMap, err := utils.GenerateJWT(admin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	token := tokenMap["access_token"]
	c.SetCookie("Adminjwt", token, 3600*24*30, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenMap})
}

// AdminHome returns welcome message
func AdminHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Admin Panel"})
}

// UserData returns paginated users
func UserData(c *gin.Context) {
	limit, err1 := strconv.Atoi(c.DefaultQuery("count", "10"))
	page, err2 := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err := errors.Join(err1, err2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	offset := (page - 1) * limit
	var users []UserResponse

	db := database.InitDB()
	if err := db.Table("users").
		Select("id, first_name, last_name, email, phone").
		Limit(limit).Offset(offset).Scan(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// ShowAllOrders returns all orders
func ShowAllOrders(c *gin.Context) {
	var orders []OrderResponse

	db := database.InitDB()
	if err := db.Table("oder_items").
		Select("order_id, user_id_no AS user_id, order_status AS status, payment_m AS payment_mode, total_amount, created_at").
		Order("user_id_no ASC").
		Scan(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// ShowOrderById returns orders by user ID
func ShowOrderById(c *gin.Context) {
	userID := c.Param("id")
	var orders []OrderResponse

	db := database.InitDB()
	if err := db.Table("oder_items").
		Where("user_id_no = ?", userID).
		Select("order_id, order_status AS status, payment_m AS payment_mode, total_amount, created_at").
		Order("order_id ASC").
		Scan(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// BlockUser sets block_status=true for a user
func BlockUser(c *gin.Context) {
	userID := c.Param("id")

	db := database.InitDB()
	if err := db.Model(&models.User{}).Where("id = ?", userID).Update("block_status", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to block user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User blocked successfully"})
}

// UnBlockUser sets block_status=false for a user
func UnBlockUser(c *gin.Context) {
	userID := c.Param("id")

	db := database.InitDB()
	if err := db.Model(&models.User{}).Where("id = ?", userID).Update("block_status", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unblock user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User unblocked successfully"})
}
