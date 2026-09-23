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

// DeleteComment godoc
// @Summary		Delete a comment
// @Description	Deletes a comment. Allowed for the comment owner or roles admin/superadmin.
// @Tags			Comments
// @Produce		json
// @Security		BearerAuth
// @Param			commentId	path		int	true	"Comment ID"
// @Success		200			{object}	map[string]interface{}	"comment deleted"
// @Failure		400			{object}	map[string]interface{}	"Invalid comment ID"
// @Failure		401			{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403			{object}	map[string]interface{}	"Not allowed to delete this comment"
// @Failure		404			{object}	map[string]interface{}	"Comment not found"
// @Failure		500			{object}	map[string]interface{}	"Failed to delete comment"
// @Router			/comments/{commentId} [delete]
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
