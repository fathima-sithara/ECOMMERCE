package routes

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	// Define API version
	apiVersion := "/api/v1"

	v1 := router.Group(apiVersion)

	UserRoutes(v1)
	// AdminRoutes(v1)
}
