package movieseriequeries

import (
	"be/internals/storage"
	"github.com/gin-gonic/gin"

	"be/internals/handler"
)

type GetMovieSeriesHandler struct {
	r storage.MovieSerieRepository
}

func NewGetMovieSeriesHandler(r storage.MovieSerieRepository) *GetMovieSeriesHandler {
	return &GetMovieSeriesHandler{r: r}
}

func (h *GetMovieSeriesHandler) Handle(c *gin.Context) {
	ms, err := h.r.GetAllMoviesAndSeries()
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get movies and series")
		return
	}
	handler.WriteSuccess(c, ms)
}
