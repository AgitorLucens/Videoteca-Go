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
