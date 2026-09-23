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

// UpdateUsername godoc
// @Summary		Update the authenticated user's username
// @Description	Changes the username of the currently authenticated user.
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		updateUsernameRequest	true	"New username (min 3 chars)"
// @Success		200		{object}	map[string]interface{}	"username updated"
// @Failure		400		{object}	map[string]interface{}	"Invalid user ID or username"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		404		{object}	map[string]interface{}	"User not found"
// @Failure		500		{object}	map[string]interface{}	"Failed to update username"
// @Router			/user/username [put]
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
