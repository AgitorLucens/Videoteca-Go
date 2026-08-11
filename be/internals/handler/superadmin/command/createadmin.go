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
