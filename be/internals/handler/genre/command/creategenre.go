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

// CreateGenre godoc
// @Summary		Create a genre (admin)
// @Description	Creates a new genre. Requires role: admin.
// @Tags			Admin - Genres
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		storage.CreateGenreRequest	true	"Genre data"
// @Success		201		{object}	map[string]interface{}		"data: created genre"
// @Failure		400		{object}	map[string]interface{}		"Invalid request payload"
// @Failure		401		{object}	map[string]interface{}		"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}		"Admin role required"
// @Failure		500		{object}	map[string]interface{}		"Failed to create genre"
// @Router			/admin/genres [post]
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
