package home

import (
	"be/internals/storage"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetMovieAndSerieHandler struct{
	r storage.HomeRepository
}

func NewLoginHandler(r storage.HomeRepository) *GetMovieAndSerieHandler{
	return &GetMovieAndSerieHandler{
		r: r,
	}
}

// GetMoviesAndSeries godoc
// @Summary		List movies and series
// @Description	Returns the full catalog of movies and series. Available to roles: user, admin.
// @Tags			Movies & Series
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	storage.HomeResponse
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Insufficient role"
// @Failure		500	{object}	map[string]interface{}	"Failed to get movies and series"
// @Router			/movieseries [get]
func (h *GetMovieAndSerieHandler) Handle(c *gin.Context){
	ms, err := h.r.GetMoviesAndSeries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get movies and series"})
		return
	}
	
	c.JSON(http.StatusOK, &storage.HomeResponse{
		Status: http.StatusText(http.StatusOK),
		MoviesAndSeries: ms,
	})
}