package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todo/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Get user ID from URL
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// Get requesting user ID from JWT
	requestingUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	// Only allow users to delete themselves or admins to delete users
	if uint(userID) != requestingUserID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "can only delete your own account"})
		return
	}

	// Delete the user
	if err := h.userService.DeleteUser(uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
    // In production, add admin role check here
    users, err := h.userService.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
        return
    }

    // Filter sensitive data before responding
    var response []map[string]interface{}
    for _, user := range users {
        response = append(response, map[string]interface{}{
            "id":        user.ID,
            "email":     user.Email,
            "created_at": user.CreatedAt,
            "is_active": user.IsActive,
        })
    }

    c.JSON(http.StatusOK, response)
}