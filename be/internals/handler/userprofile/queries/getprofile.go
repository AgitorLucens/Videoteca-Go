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

// GetProfile godoc
// @Summary		Get current user's profile
// @Description	Returns the profile of the authenticated user, including a base64 data-URI profile picture when set.
// @Tags			User Profile
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	map[string]interface{}	"id, username, email, name, profile_picture"
// @Failure		400	{object}	map[string]interface{}	"Invalid user ID"
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		404	{object}	map[string]interface{}	"User not found"
// @Router			/user/profile [get]
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
