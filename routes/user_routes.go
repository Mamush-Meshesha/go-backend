package routes

import (
	"github.com/gin-gonic/gin"
	"todo/handlers"
	"todo/middleware"
	"todo/services"
)

func SetupUserRoutes(router *gin.Engine, userService *services.UserService) {
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}