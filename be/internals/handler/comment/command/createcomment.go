package comment

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateCommentHandler struct {
	r storage.MovieSerieRepository
}

func NewCreateCommentHandler(r storage.MovieSerieRepository) *CreateCommentHandler {
	return &CreateCommentHandler{r: r}
}

type createCommentRequest struct {
	Comment string `json:"comment" binding:"required"`
}

func (h *CreateCommentHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req createCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment is required"})
		return
	}

	email, _ := c.Get("email")
	userEmail, _ := email.(string)

	if err := h.r.CreateComment(uint(id), userEmail, req.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "comment created"})
}
