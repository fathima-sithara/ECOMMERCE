package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/models"
)

func AddToCart(c *gin.Context) {
	type AddToCartRequest struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  uint `json:"quantity" binding:"required,min=1"`
	}

	var req AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	db := database.InitDB()

	// Check if product exists
	var product models.Product
	if err := db.First(&product, req.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Parse price from string to uint
	priceUint, err := strconv.ParseUint(product.Price, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid product price format"})
		return
	}

	// Check stock
	if req.Quantity > product.Stock {
		c.JSON(http.StatusConflict, gin.H{"error": "Insufficient stock"})
		return
	}

	var cart models.Cart
	err = db.Where("product_id = ? AND user_id = ?", req.ProductID, userID).First(&cart).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new cart item
		newCart := models.Cart{
			ProductId:  req.ProductID,
			Quantity:   req.Quantity,
			Price:      uint(priceUint),
			TotalPrice: uint(priceUint) * req.Quantity,
			UserId:     uint(userID),
		}
		if err := db.Create(&newCart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query cart"})
		return
	}

	// Update existing cart item
	cart.Quantity += req.Quantity
	cart.TotalPrice = cart.Quantity * uint(priceUint)

	if err := db.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully"})
}

type CartItemResponse struct {
	ProductName string `json:"product_name"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	TotalPrice  uint   `json:"total_price"`
}

func ViewCart(c *gin.Context) {
	// Get user ID from token (set in middleware)
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in token"})
		return
	}

	var cartItems []CartItemResponse
	db := database.InitDB()

	// Join carts and products
	err = db.Table("carts").
		Select("products.product_name, carts.quantity, carts.price, carts.total_price").
		Joins("INNER JOIN products ON products.id = carts.product_id").
		Where("carts.user_id = ?", userID).
		Scan(&cartItems).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items", "details": err.Error()})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Cart is empty"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_items": cartItems,
	})
}

func DeleteCart(c *gin.Context) {
	// Parse cart ID from path
	cartIDParam := c.Param("id")
	cartID, err := strconv.Atoi(cartIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	// Get user ID from JWT context
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	db := database.InitDB()

	// Check if the cart item exists for this user
	var cart models.Cart
	if err := db.Where("id = ? AND user_id = ?", cartID, userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found or unauthorized"})
		return
	}

	// Perform the deletion
	if err := db.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item deleted successfully"})
}
