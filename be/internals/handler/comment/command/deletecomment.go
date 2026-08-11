package comment

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCommentHandler struct {
	r storage.MovieSerieRepository
}

func NewDeleteCommentHandler(r storage.MovieSerieRepository) *DeleteCommentHandler {
	return &DeleteCommentHandler{r: r}
}

func (h *DeleteCommentHandler) Handle(c *gin.Context) {
	commentIDStr := c.Param("commentId")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment ID"})
		return
	}

	comment, err := h.r.GetCommentByID(uint(commentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	email, _ := c.Get("email")
	userEmail, _ := email.(string)

	role, _ := c.Get("role")
	userRole, _ := role.(string)

	if comment.AppUser != userEmail && userRole != "admin" && userRole != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "not allowed to delete this comment"})
		return
	}

	if err := h.r.DeleteComment(uint(commentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}
