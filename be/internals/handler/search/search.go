package search

import (
	"be/internals/storage"
	"github.com/gin-gonic/gin"

	"be/internals/handler"
)

type SearchMoviesHandler struct {
	r storage.SearchRepository
}

func NewSearchMoviesHandler(r storage.SearchRepository) *SearchMoviesHandler {
	return &SearchMoviesHandler{r: r}
}

// Search godoc
// @Summary		Search movies, series and actors
// @Description	Searches movies/series by title and actors by name, returning a combined result list.
// @Tags			Search
// @Produce		json
// @Security		BearerAuth
// @Param			q	query		string	true	"Search query"
// @Success		200	{object}	handler.Response{data=[]storage.SearchResult}
// @Failure		400	{object}	handler.Response	"Search query is required"
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		500	{object}	handler.Response	"Failed to search"
// @Router			/search [get]
func (h *SearchMoviesHandler) Handle(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		handler.WriteBadRequest(c, "search query is required")
		return
	}

	ms, err := h.r.SearchMovies(query)
	if err != nil {
		handler.WriteInternalServerError(c, "failed to search movies")
		return
	}

	actors, err := h.r.SearchActors(query)
	if err != nil {
		handler.WriteInternalServerError(c, "failed to search actors")
		return
	}

	var results []storage.SearchResult

	for _, m := range ms {
		msType := "movie"
		if m.MSType != "" {
			msType = m.MSType
		}
		results = append(results, storage.SearchResult{
			ID:     m.ID,
			Title:  m.Title,
			Type:   msType,
			MSType: msType,
		})
	}

	for _, a := range actors {
		title := a.ActorFirstName + " " + a.ActorLastName
		photo := ""
		if a.ActorPhoto != nil {
			photo = *a.ActorPhoto
		}
		results = append(results, storage.SearchResult{
			ID:    a.ActorID,
			Title: title,
			Type:  "actor",
			Photo: photo,
		})
	}

	handler.WriteSuccess(c, results)
}
