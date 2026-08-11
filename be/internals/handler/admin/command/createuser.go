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