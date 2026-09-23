package superadminqueries

import (
	"be/internals/rbac"
	"github.com/gin-gonic/gin"

	"be/internals/handler"
)

type GetAdminsHandler struct {
	r rbac.UserRepository
}

func NewGetAdminsHandler(r rbac.UserRepository) *GetAdminsHandler {
	return &GetAdminsHandler{r: r}
}

// GetAdmins godoc
// @Summary		List all administrators (superadmin)
// @Description	Returns all users with the admin role. Requires role: superadmin.
// @Tags			Superadmin
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	handler.Response{data=[]rbac.User}
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Superadmin role required"
// @Failure		500	{object}	handler.Response	"Failed to get admins"
// @Router			/superadmin/admins [get]
func (h *GetAdminsHandler) Handle(c *gin.Context) {
	users, err := h.r.GetUsersByRole("admin")
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get admins")
		return
	}
	handler.WriteSuccess(c, users)
}
