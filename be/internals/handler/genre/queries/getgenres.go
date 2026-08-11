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

func (h *GetGenresHandler) Handle(c *gin.Context) {
	genres, err := h.r.GetAllGenres()
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get genres")
		return
	}
	handler.WriteSuccess(c, genres)
}
