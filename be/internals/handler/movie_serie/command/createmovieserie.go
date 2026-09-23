package movieseriecommand

import (
	"be/internals/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMovieSerieHandler struct {
	r storage.MovieSerieRepository
}

func NewCreateMovieSerieHandler(r storage.MovieSerieRepository) *CreateMovieSerieHandler {
	return &CreateMovieSerieHandler{r: r}
}

// CreateMovieSerie godoc
// @Summary		Create a movie or series (admin)
// @Description	Creates a new movie/series optionally linked to genres and actors. Requires role: admin.
// @Tags			Admin - Movies & Series
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		storage.CreateMovieSerieRequest	true	"Movie/Series data"
// @Success		201		{object}	map[string]interface{}			"data: created movie/series"
// @Failure		400		{object}	map[string]interface{}			"Invalid request payload"
// @Failure		401		{object}	map[string]interface{}			"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}			"Admin role required"
// @Failure		500		{object}	map[string]interface{}			"Failed to create movie/serie"
// @Router			/admin/movieseries [post]
func (h *CreateMovieSerieHandler) Handle(c *gin.Context) {
	var req storage.CreateMovieSerieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ms, err := h.r.CreateMovieSerie(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create movie/serie"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": ms})
}
