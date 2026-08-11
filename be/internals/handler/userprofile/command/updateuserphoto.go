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
