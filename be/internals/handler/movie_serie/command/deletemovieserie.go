package movieseriecommand

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMovieSerieHandler struct {
	r storage.MovieSerieRepository
}

func NewDeleteMovieSerieHandler(r storage.MovieSerieRepository) *DeleteMovieSerieHandler {
	return &DeleteMovieSerieHandler{r: r}
}

// DeleteMovieSerie godoc
// @Summary		Delete a movie or series (admin)
// @Description	Deletes a movie/series by ID. Requires role: admin.
// @Tags			Admin - Movies & Series
// @Produce		json
// @Security		BearerAuth
// @Param			id	path		int	true	"Movie/Series ID"
// @Success		200	{object}	map[string]interface{}	"movie/serie deleted successfully"
// @Failure		400	{object}	map[string]interface{}	"Invalid ID"
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Admin role required"
// @Failure		500	{object}	map[string]interface{}	"Failed to delete movie/serie"
// @Router			/admin/movieseries/{id} [delete]
func (h *DeleteMovieSerieHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.r.DeleteMovieSerie(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete movie/serie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "movie/serie deleted successfully"})
}
