package genrecommand

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteGenreHandler struct {
	r storage.GenreRepository
}

func NewDeleteGenreHandler(r storage.GenreRepository) *DeleteGenreHandler {
	return &DeleteGenreHandler{r: r}
}

func (h *DeleteGenreHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.r.DeleteGenre(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "genre deleted successfully"})
}
