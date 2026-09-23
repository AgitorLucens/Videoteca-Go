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

// ForgotPassword godoc
// @Summary		Check account existence for password recovery
// @Description	Looks up an account by email and returns the user ID if found.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			email	body		ForgotPasswordRequest	true	"Account email"
// @Success		200		{object}	map[string]interface{}	"message and userID"
// @Failure		400		{object}	map[string]interface{}	"Valid email is required"
// @Failure		404		{object}	map[string]interface{}	"No account found with that email"
// @Router			/forgot-password [post]
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
