package genrequeries

import (
	"be/internals/storage"
	"github.com/gin-gonic/gin"

	"be/internals/handler"
)

type GetGenresHandler struct {
	r storage.GenreRepository
}

func NewGetGenresHandler(r storage.GenreRepository) *GetGenresHandler {
	return &GetGenresHandler{r: r}
}

// GetGenres godoc
// @Summary		List all genres (admin)
// @Description	Returns all available genres. Requires role: admin.
// @Tags			Admin - Genres
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	handler.Response{data=[]storage.Genre}
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Admin role required"
// @Failure		500	{object}	handler.Response	"Failed to get genres"
// @Router			/admin/genres [get]
func (h *GetGenresHandler) Handle(c *gin.Context) {
	genres, err := h.r.GetAllGenres()
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get genres")
		return
	}
	handler.WriteSuccess(c, genres)
}
