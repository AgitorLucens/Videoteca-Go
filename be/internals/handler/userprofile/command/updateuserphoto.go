package userprofile

import (
	"be/internals/rbac"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UpdatePhotoHandler struct {
	repo *rbac.Repository
}

func NewUpdatePhotoHandler(repo *rbac.Repository) *UpdatePhotoHandler {
	return &UpdatePhotoHandler{repo: repo}
}

type updatePhotoRequest struct {
	Photo string `json:"photo" binding:"required"`
}

// UpdatePhoto godoc
// @Summary		Update the authenticated user's profile picture
// @Description	Uploads a new profile picture as a base64-encoded image (optionally with a data-URI prefix).
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		updatePhotoRequest		true	"Base64 image data"
// @Success		200		{object}	map[string]interface{}	"photo updated"
// @Failure		400		{object}	map[string]interface{}	"Invalid user ID or image data"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		500		{object}	map[string]interface{}	"Failed to update photo"
// @Router			/user/picture [put]
func (h *UpdatePhotoHandler) Handle(c *gin.Context) {
	userIDStr, _ := c.Get("userID")
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var req updatePhotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "photo data is required"})
		return
	}

	data := req.Photo
	if strings.Contains(data, ",") {
		data = strings.SplitN(data, ",", 2)[1]
	}

	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image data"})
		return
	}

	if err := h.repo.UpdateUserProfilePicture(uint(userID), bytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "photo updated"})
}
