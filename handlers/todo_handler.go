package handlers

import (
	"net/http"
	"strconv"
	"todo/models"
	"todo/services"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler (service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	todo,err := h.service.GetTodoByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todos not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
    // Get user ID from JWT
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
        return
    }

    // Parse todo ID from URL
    todoID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo ID"})
        return
    }

    // Fetch existing todo
    existingTodo, err := h.service.GetTodoByID(uint(todoID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
        return
    }

    // Verify ownership
    if existingTodo.UserID != userID.(uint) {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to update this todo"})
        return
    }

    // Parse update data
    var updateData models.Todo
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Apply updates (only to allowed fields)
    existingTodo.Title = updateData.Title
    existingTodo.Description = updateData.Description
    existingTodo.Completed = updateData.Completed

    // Save changes
    if err := h.service.UpdateTodo(existingTodo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "todo updated successfully",
        "todo":    existingTodo,
    })
}

func (h *TodoHandler) DeleteTodo(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalide id"})
		return
	}
	if err := h.service.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}