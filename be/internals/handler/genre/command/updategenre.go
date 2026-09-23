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

// UpdateGenre godoc
// @Summary		Update a genre (admin)
// @Description	Updates an existing genre. Requires role: admin.
// @Tags			Admin - Genres
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			id		path		int							true	"Genre ID"
// @Param			request	body		storage.UpdateGenreRequest	true	"Genre data"
// @Success		200		{object}	map[string]interface{}		"data: updated genre"
// @Failure		400		{object}	map[string]interface{}		"Invalid ID or payload"
// @Failure		401		{object}	map[string]interface{}		"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}		"Admin role required"
// @Failure		500		{object}	map[string]interface{}		"Failed to update genre"
// @Router			/admin/genres/{id} [put]
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
