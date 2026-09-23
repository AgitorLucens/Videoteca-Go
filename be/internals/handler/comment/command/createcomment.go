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

// CreateComment godoc
// @Summary		Add a comment to a movie/series
// @Description	Creates a comment on behalf of the authenticated user. Available to roles: user, admin.
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			id		path		int						true	"Movie/Series ID"
// @Param			request	body		createCommentRequest	true	"Comment text"
// @Success		201		{object}	map[string]interface{}	"comment created"
// @Failure		400		{object}	map[string]interface{}	"Invalid ID or missing comment"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		500		{object}	map[string]interface{}	"Failed to create comment"
// @Router			/movieseries/{id}/comments [post]
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
