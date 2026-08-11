package userprofile

import (
	"be/internals/rbac"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEmailHandler struct {
	repo *rbac.Repository
}

func NewUpdateEmailHandler(repo *rbac.Repository) *UpdateEmailHandler {
	return &UpdateEmailHandler{repo: repo}
}

type updateEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (h *UpdateEmailHandler) Handle(c *gin.Context) {
	userIDStr, _ := c.Get("userID")
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var req updateEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "valid email is required"})
		return
	}

	u, err := h.repo.GetUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := h.repo.UpdateUser(u.ID, u.Name, u.UserName, req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email updated"})
}
