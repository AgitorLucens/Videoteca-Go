package rating

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateRatingHandler struct {
	r storage.MovieSerieRepository
}

func NewCreateRatingHandler(r storage.MovieSerieRepository) *CreateRatingHandler {
	return &CreateRatingHandler{r: r}
}

type createRatingRequest struct {
	Rating int `json:"rating" binding:"required"`
}

// RateMovieSerie godoc
// @Summary		Rate a movie or series
// @Description	Creates or updates the authenticated user's rating (1-5) for a movie/series and returns updated aggregate rating data. Available to roles: user, admin.
// @Tags			Ratings
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			id		path		int						true	"Movie/Series ID"
// @Param			request	body		createRatingRequest		true	"Rating value (1-5)"
// @Success		200		{object}	map[string]interface{}	"ratingData"
// @Failure		400		{object}	map[string]interface{}	"Invalid ID or rating out of range"
// @Failure		401		{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		500		{object}	map[string]interface{}	"Failed to save or fetch rating"
// @Router			/movieseries/{id}/rate [post]
func (h *CreateRatingHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req createRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Rating < 1 || req.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rating must be between 1 and 5"})
		return
	}

	email, _ := c.Get("email")
	userEmail, _ := email.(string)

	if err := h.r.UpsertRating(uint(id), userEmail, req.Rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save rating"})
		return
	}

	ratingData, err := h.r.GetRatingDataByMsID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch updated rating"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ratingData": ratingData})
}
