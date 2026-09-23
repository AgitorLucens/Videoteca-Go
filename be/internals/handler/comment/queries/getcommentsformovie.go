package comment

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCommentsForMovieHandler struct {
	r storage.MovieSerieRepository
}

func NewGetCommentsForMovieHandler(r storage.MovieSerieRepository) *GetCommentsForMovieHandler {
	return &GetCommentsForMovieHandler{r: r}
}

// GetCommentsForMovie godoc
// @Summary		List comments for a movie/series
// @Description	Returns a paginated list of comments for the given movie/series. Available to roles: user, admin.
// @Tags			Comments
// @Produce		json
// @Security		BearerAuth
// @Param			id		path		int		true	"Movie/Series ID"
// @Param			page	query		int		false	"Page number"	default(1)
// @Param			limit	query		int		false	"Items per page (max 50)"	default(10)
// @Success		200		{object}	map[string]interface{}	"comments, total, page, limit, totalPages"
// @Failure		400		{object}	map[string]interface{}	"Invalid ID"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		500		{object}	map[string]interface{}	"Failed to fetch comments"
// @Router			/movieseries/{id}/comments [get]
func (h *GetCommentsForMovieHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}

	comments, total, err := h.r.GetCommentsByMsIDPaginated(uint(id), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch comments"})
		return
	}

	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, gin.H{
		"comments":    comments,
		"total":       total,
		"page":        page,
		"limit":       limit,
		"totalPages":  totalPages,
	})
}
