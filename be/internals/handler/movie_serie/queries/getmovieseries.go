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

// GetMovieSeries godoc
// @Summary		List all movies and series (admin)
// @Description	Returns the full catalog for administration. Requires role: admin.
// @Tags			Admin - Movies & Series
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	handler.Response{data=[]storage.MovieSerie}
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Admin role required"
// @Failure		500	{object}	handler.Response	"Failed to get movies and series"
// @Router			/admin/movieseries [get]
func (h *GetMovieSeriesHandler) Handle(c *gin.Context) {
	ms, err := h.r.GetAllMoviesAndSeries()
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get movies and series")
		return
	}
	handler.WriteSuccess(c, ms)
}
