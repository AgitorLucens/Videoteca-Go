package admin

import (
	"be/internals/rbac"
	//"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type PostUserHandler struct{
	r *rbac.Repository
}

func NewPostUserHandler(r *rbac.Repository) *PostUserHandler{
	return &PostUserHandler{
		r: r,
	}
}

// RegisterUser godoc
// @Summary		Register a new user
// @Description	Public registration endpoint. Creates a user with the given role.
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			request	body		rbac.CreateUserRequest	true	"User registration data"
// @Success		201		{object}	map[string]interface{}	"msg and user"
// @Failure		400		{object}	map[string]interface{}	"Invalid request payload"
// @Failure		500		{object}	map[string]interface{}	"Failed to create user"
// @Router			/users [post]
func (h *PostUserHandler) Handle(c *gin.Context){
	var req rbac.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Received request to create user:", req)
	_, err := h.r.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "user created successfully",
		"user":    req,})
}