package movieseriecommand

import (
	"be/internals/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMovieSerieHandler struct {
	r storage.MovieSerieRepository
}

func NewCreateMovieSerieHandler(r storage.MovieSerieRepository) *CreateMovieSerieHandler {
	return &CreateMovieSerieHandler{r: r}
}

func (h *CreateMovieSerieHandler) Handle(c *gin.Context) {
	var req storage.CreateMovieSerieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ms, err := h.r.CreateMovieSerie(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create movie/serie"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": ms})
}
