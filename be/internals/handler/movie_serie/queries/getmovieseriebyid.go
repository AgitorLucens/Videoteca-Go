package movieseriequeries

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMovieSerieByIDHandler struct {
	r storage.MovieSerieRepository
}

func NewGetMovieSerieByIDHandler(r storage.MovieSerieRepository) *GetMovieSerieByIDHandler {
	return &GetMovieSerieByIDHandler{r: r}
}

// GetMovieSerieByID godoc
// @Summary		Get movie/series details
// @Description	Returns details of a movie/series: metadata, actors, genres, rating data, episodes and the authenticated user's own rating. Available to roles: user, admin.
// @Tags			Movies & Series
// @Produce		json
// @Security		BearerAuth
// @Param			id	path		int	true	"Movie/Series ID"
// @Success		200	{object}	map[string]interface{}	"movieSerie, actors, genres, ratingData, episodes, genreIds, actorIds, userRating"
// @Failure		400	{object}	map[string]interface{}	"Invalid ID"
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		404	{object}	map[string]interface{}	"Movie/serie not found"
// @Router			/movieseries/{id} [get]
func (h *GetMovieSerieByIDHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	ms, err := h.r.GetMovieSerieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie/serie not found"})
		return
	}

	genres, _ := h.r.GetGenresByMsID(uint(id))
	actors, _ := h.r.GetActorsByMsID(uint(id))
	episodes, _ := h.r.GetEpisodesByMsID(uint(id))
	ratingData, _ := h.r.GetRatingDataByMsID(uint(id))

	userRating := 0
	email, _ := c.Get("email")
	userEmail, _ := email.(string)
	if userEmail != "" {
		ur, err := h.r.GetUserRating(uint(id), userEmail)
		if err == nil {
			userRating = ur
		}
	}

	genreNames := ""
	genreIDs := []uint{}
	for i, g := range genres {
		if i > 0 {
			genreNames += ", "
		}
		genreNames += g.GenreName
		genreIDs = append(genreIDs, g.GenreID)
	}

	actorIDs := []uint{}
	for _, a := range actors {
		actorIDs = append(actorIDs, a.ActorID)
	}

	data := storage.MovieOrSerieData{
		MovieSerie: *ms,
		Actors:     actors,
		Genres:     genreNames,
		RatingData: ratingData,
		Episodes:   episodes,
	}

	response := gin.H{
		"movieSerie": data.MovieSerie,
		"actors":     data.Actors,
		"genres":     data.Genres,
		"ratingData": data.RatingData,
		"episodes":   data.Episodes,
		"genreIds":   genreIDs,
		"actorIds":   actorIDs,
		"userRating": userRating,
	}

	c.JSON(http.StatusOK, response)
}
