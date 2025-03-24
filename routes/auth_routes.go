package routes

import (
	"todo/handlers"
	"todo/services"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authService *services.AuthService) {
	authHandler := handlers.NewAuthHandler(authService)

	authRoutes := router.Group("/auth")

	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}
}