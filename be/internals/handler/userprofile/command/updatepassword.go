package userprofile

import (
	"be/internals/rbac"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordHandler struct {
	repo *rbac.Repository
}

func NewUpdatePasswordHandler(repo *rbac.Repository) *UpdatePasswordHandler {
	return &UpdatePasswordHandler{repo: repo}
}

type updatePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
}

// UpdatePassword godoc
// @Summary		Update the authenticated user's password
// @Description	Changes the password of the currently authenticated user; the current password must match.
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		updatePasswordRequest	true	"Current and new password"
// @Success		200		{object}	map[string]interface{}	"password updated"
// @Failure		400		{object}	map[string]interface{}	"Invalid payload or current password incorrect"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		500		{object}	map[string]interface{}	"Failed to update password"
// @Router			/user/password [put]
func (h *UpdatePasswordHandler) Handle(c *gin.Context) {
	userIDStr, _ := c.Get("userID")
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var req updatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "current password and new password (min 8 chars) are required"})
		return
	}

	if err := h.repo.UpdateUserPassword(uint(userID), req.CurrentPassword, req.NewPassword); err != nil {
		if err.Error() == "current password is incorrect" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "current password is incorrect"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated"})
}
