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