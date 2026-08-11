package admin

import (
	"be/internals/rbac"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetUserHandler struct{
	rep *rbac.Repository
}

func NewGetUserHandler(r *rbac.Repository) *GetUserHandler{
	return &GetUserHandler{
		rep: r,
	}
}

func (h *GetUserHandler) Handle(c *gin.Context){
	idStr := c.Param("id")
	id, err  := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	u, err := h.rep.GetUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}