package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.PUT("/profile", controllers.UpdateProfile)
		auth.PUT("/change-password", controllers.ChangePassword)
	}

	return r
}
