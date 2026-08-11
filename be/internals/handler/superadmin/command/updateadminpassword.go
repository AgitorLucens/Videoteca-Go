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
