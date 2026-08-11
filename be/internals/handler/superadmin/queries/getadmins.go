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

func (h *GetAdminsHandler) Handle(c *gin.Context) {
	users, err := h.r.GetUsersByRole("admin")
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get admins")
		return
	}
	handler.WriteSuccess(c, users)
}
