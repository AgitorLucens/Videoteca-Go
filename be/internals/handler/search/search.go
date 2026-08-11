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
