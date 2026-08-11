package genrecommand

import (
	"be/internals/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGenreHandler struct {
	r storage.GenreRepository
}

func NewCreateGenreHandler(r storage.GenreRepository) *CreateGenreHandler {
	return &CreateGenreHandler{r: r}
}

func (h *CreateGenreHandler) Handle(c *gin.Context) {
	var req storage.CreateGenreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre, err := h.r.CreateGenre(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create genre"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": genre})
}
