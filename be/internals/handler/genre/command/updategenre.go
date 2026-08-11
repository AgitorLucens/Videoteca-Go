package genrecommand

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateGenreHandler struct {
	r storage.GenreRepository
}

func NewUpdateGenreHandler(r storage.GenreRepository) *UpdateGenreHandler {
	return &UpdateGenreHandler{r: r}
}

func (h *UpdateGenreHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req storage.UpdateGenreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre, err := h.r.UpdateGenre(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}
