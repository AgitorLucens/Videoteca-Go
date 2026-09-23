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

// DeleteGenre godoc
// @Summary		Delete a genre (admin)
// @Description	Deletes a genre by ID. Requires role: admin.
// @Tags			Admin - Genres
// @Produce		json
// @Security		BearerAuth
// @Param			id	path		int	true	"Genre ID"
// @Success		200	{object}	map[string]interface{}	"genre deleted successfully"
// @Failure		400	{object}	map[string]interface{}	"Invalid ID"
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Admin role required"
// @Failure		500	{object}	map[string]interface{}	"Failed to delete genre"
// @Router			/admin/genres/{id} [delete]
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
