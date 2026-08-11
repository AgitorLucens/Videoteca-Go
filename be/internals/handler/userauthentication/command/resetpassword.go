package auserauthentication

import (
	"be/internals/rbac"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResetPasswordHandler struct {
	r *rbac.Repository
}

func NewResetPasswordHandler(r *rbac.Repository) *ResetPasswordHandler {
	return &ResetPasswordHandler{r: r}
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}

func (h *ResetPasswordHandler) Handle(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "valid email and password (min 6 chars) are required"})
		return
	}

	user, err := h.r.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no account found with that email"})
		return
	}

	if err := h.r.SetNewPassword(user.ID, req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reset password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}
