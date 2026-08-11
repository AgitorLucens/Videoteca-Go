package userprofile

import (
	"be/internals/rbac"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProfileHandler struct {
	repo *rbac.Repository
}

func NewGetProfileHandler(repo *rbac.Repository) *GetProfileHandler {
	return &GetProfileHandler{repo: repo}
}

func (h *GetProfileHandler) Handle(c *gin.Context) {
	userIDStr, _ := c.Get("userID")
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	u, err := h.repo.GetUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var photo string
	if len(u.ProfilePicture) > 0 {
		photo = "data:image/png;base64," + base64.StdEncoding.EncodeToString(u.ProfilePicture)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             u.ID,
		"username":       u.UserName,
		"email":          u.Email,
		"name":           u.Name,
		"profile_picture": photo,
	})
}
