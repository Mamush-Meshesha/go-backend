package routes

import (
	"todo/handlers"
	"todo/services"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine, todoService *services.TodoService){
	todoHandler := handlers.NewTodoHandler(todoService)
	
	todoRoutes := router.Group("/todos")

	{
		todoRoutes.POST("/", todoHandler.CreateTodo)
		todoRoutes.GET("/", todoHandler.GetAllTodos)
		todoRoutes.GET("/:id", todoHandler.GetTodoByID)
		todoRoutes.PUT("/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
	}
}