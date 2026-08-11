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
