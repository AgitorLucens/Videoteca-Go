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

// UpdateEmail godoc
// @Summary		Update the authenticated user's email
// @Description	Changes the email of the currently authenticated user.
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		updateEmailRequest		true	"New email"
// @Success		200		{object}	map[string]interface{}	"email updated"
// @Failure		400		{object}	map[string]interface{}	"Invalid user ID or email"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		404		{object}	map[string]interface{}	"User not found"
// @Failure		500		{object}	map[string]interface{}	"Failed to update email"
// @Router			/user/email [put]
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
