package movieseriecommand

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMovieSerieHandler struct {
	r storage.MovieSerieRepository
}

func NewUpdateMovieSerieHandler(r storage.MovieSerieRepository) *UpdateMovieSerieHandler {
	return &UpdateMovieSerieHandler{r: r}
}

// UpdateMovieSerie godoc
// @Summary		Update a movie or series (admin)
// @Description	Updates an existing movie/series, optionally replacing linked genres and actors. Requires role: admin.
// @Tags			Admin - Movies & Series
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			id		path		int								true	"Movie/Series ID"
// @Param			request	body		storage.UpdateMovieSerieRequest	true	"Fields to update"
// @Success		200		{object}	map[string]interface{}			"data: updated movie/series"
// @Failure		400		{object}	map[string]interface{}			"Invalid ID or payload"
// @Failure		401		{object}	map[string]interface{}			"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}			"Admin role required"
// @Failure		500		{object}	map[string]interface{}			"Failed to update movie/serie"
// @Router			/admin/movieseries/{id} [put]
func (h *UpdateMovieSerieHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req storage.UpdateMovieSerieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ms, err := h.r.UpdateMovieSerie(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update movie/serie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ms})
}
