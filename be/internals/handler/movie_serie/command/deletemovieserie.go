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
