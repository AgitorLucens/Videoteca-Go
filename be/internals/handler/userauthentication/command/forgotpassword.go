package auserauthentication

import (
	"be/internals/rbac"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	r *rbac.Repository
}

func NewForgotPasswordHandler(r *rbac.Repository) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{r: r}
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (h *ForgotPasswordHandler) Handle(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "valid email is required"})
		return
	}

	user, err := h.r.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no account found with that email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account found", "userID": user.ID})
}
