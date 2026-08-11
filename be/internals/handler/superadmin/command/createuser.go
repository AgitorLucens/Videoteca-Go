package superadmincommand

import (
	"be/internals/rbac"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	r *rbac.Repository
}

func NewCreateUserHandler(r *rbac.Repository) *CreateUserHandler {
	return &CreateUserHandler{
		r: r,
	}
}

func (h *CreateUserHandler) Handle(c *gin.Context) {
	var req rbac.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.r.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"user":    user,
	})
}
