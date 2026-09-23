package superadmincommand

import (
	"be/internals/rbac"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAdminHandler struct {
	r rbac.UserRepository
}

func NewCreateAdminHandler(r rbac.UserRepository) *CreateAdminHandler {
	return &CreateAdminHandler{r: r}
}

// CreateAdmin godoc
// @Summary		Create an administrator (superadmin)
// @Description	Creates a new user account with the admin role forced. Requires role: superadmin.
// @Tags			Superadmin
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		rbac.CreateUserRequest	true	"Admin account data (role is forced to admin)"
// @Success		201		{object}	map[string]interface{}	"data: created admin user"
// @Failure		400		{object}	map[string]interface{}	"Invalid request payload"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}	"Superadmin role required"
// @Failure		500		{object}	map[string]interface{}	"Failed to create admin"
// @Router			/superadmin/admins [post]
func (h *CreateAdminHandler) Handle(c *gin.Context) {
	var req rbac.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Role = "admin"

	user, err := h.r.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
