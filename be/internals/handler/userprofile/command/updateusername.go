package userprofile

import (
	"be/internals/rbac"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUsernameHandler struct {
	repo *rbac.Repository
}

func NewUpdateUsernameHandler(repo *rbac.Repository) *UpdateUsernameHandler {
	return &UpdateUsernameHandler{repo: repo}
}

type updateUsernameRequest struct {
	Username string `json:"username" binding:"required,min=3"`
}

func (h *UpdateUsernameHandler) Handle(c *gin.Context) {
	userIDStr, _ := c.Get("userID")
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var req updateUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required (min 3 chars)"})
		return
	}

	u, err := h.repo.GetUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := h.repo.UpdateUser(u.ID, u.Name, req.Username, u.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update username"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "username updated"})
}
