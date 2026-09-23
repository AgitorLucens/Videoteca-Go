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

// GetUser godoc
// @Summary		Get a user by ID
// @Description	Returns the user with the given numeric ID.
// @Tags			Users
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	rbac.User
// @Failure		400	{object}	map[string]interface{}	"Invalid ID or user not found"
// @Router			/users/{id} [get]
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