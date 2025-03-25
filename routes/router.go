package routes

import (
	"todo/handlers"
	"todo/middleware"
	"todo/services"

	"github.com/gin-gonic/gin"
)

type Services struct {
	AuthService *services.AuthService
	TodoService *services.TodoService
	UserService *services.UserService
}

// Remove this - it's duplicate (already defined in handlers package)
// type UserHandler struct {
//     userService *services.UserService
// }

func SetupAllRoutes(router *gin.Engine, s Services) {
	// Auth routes (no auth middleware)
	authRoutes := router.Group("/auth")
	{
		authHandler := handlers.NewAuthHandler(s.AuthService) // Use handlers package directly
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	// Authenticated routes
	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware()) // Add middleware package reference
	{
		// Todo routes
		todoHandler := handlers.NewTodoHandler(s.TodoService)
		todoRoutes := authGroup.Group("/todos")
		{
			todoRoutes.POST("/", todoHandler.CreateTodo)
			todoRoutes.GET("/", todoHandler.GetAllTodos)
			todoRoutes.GET("/:id", todoHandler.GetTodoByID)
			todoRoutes.PUT("/:id", todoHandler.UpdateTodo)
			todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
		}

		// User routes
		userHandler := handlers.NewUserHandler(s.UserService)
		userRoutes := authGroup.Group("/users")
		{
			userRoutes.GET("/", userHandler.GetAllUsers)
			userRoutes.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}