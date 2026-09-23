package superadmincommand

import (
	"be/internals/rbac"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateAdminPasswordHandler struct {
	r rbac.UserRepository
}

func NewUpdateAdminPasswordHandler(r rbac.UserRepository) *UpdateAdminPasswordHandler {
	return &UpdateAdminPasswordHandler{r: r}
}

type UpdateAdminPasswordRequest struct {
	UserID      uint   `json:"userId" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}

// UpdateAdminPassword godoc
// @Summary		Reset an admin's password (superadmin)
// @Description	Sets a new password for the user identified by userId. Requires role: superadmin.
// @Tags			Superadmin
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		UpdateAdminPasswordRequest	true	"User ID and new password (min 6 chars)"
// @Success		200		{object}	map[string]interface{}		"password updated successfully"
// @Failure		400		{object}	map[string]interface{}		"Valid userId and password are required"
// @Failure		401		{object}	map[string]interface{}		"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}		"Superadmin role required"
// @Failure		404		{object}	map[string]interface{}		"User not found"
// @Failure		500		{object}	map[string]interface{}		"Failed to update password"
// @Router			/superadmin/admins/password [put]
func (h *UpdateAdminPasswordHandler) Handle(c *gin.Context) {
	var req UpdateAdminPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "valid userId and password (min 6 chars) are required"})
		return
	}

	user, err := h.r.GetUserById(req.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := h.r.SetUserPassword(user.ID, req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}
