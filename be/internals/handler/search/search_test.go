package search

import (
	"be/internals/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockSearchRepo struct {
	movies    []storage.MovieSerie
	actors    []storage.Actor
	searchErr error
}

func (m *mockSearchRepo) SearchMovies(query string) ([]storage.MovieSerie, error) {
	if m.searchErr != nil {
		return nil, m.searchErr
	}
	return m.movies, nil
}

func (m *mockSearchRepo) SearchActors(query string) ([]storage.Actor, error) {
	if m.searchErr != nil {
		return nil, m.searchErr
	}
	return m.actors, nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestSearchMoviesHandler_Success(t *testing.T) {
	mock := &mockSearchRepo{
		movies: []storage.MovieSerie{
			{ID: 1, Title: "Inception", MSType: "movie"},
			{ID: 2, Title: "Interstellar", MSType: "movie"},
		},
	}

	r := setupRouter()
	handler := NewSearchMoviesHandler(mock)
	r.GET("/search", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?q=Incep", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]storage.SearchResult
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["data"], 2)
	assert.Equal(t, "Inception", response["data"][0].Title)
}

func TestSearchMoviesHandler_EmptyQuery(t *testing.T) {
	mock := &mockSearchRepo{}

	r := setupRouter()
	handler := NewSearchMoviesHandler(mock)
	r.GET("/search", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?q=", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSearchMoviesHandler_MissingQuery(t *testing.T) {
	mock := &mockSearchRepo{}

	r := setupRouter()
	handler := NewSearchMoviesHandler(mock)
	r.GET("/search", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSearchMoviesHandler_RepoError(t *testing.T) {
	mock := &mockSearchRepo{searchErr: assert.AnError}

	r := setupRouter()
	handler := NewSearchMoviesHandler(mock)
	r.GET("/search", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?q=test", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSearchMoviesHandler_NoResults(t *testing.T) {
	mock := &mockSearchRepo{
		movies: []storage.MovieSerie{},
		actors: []storage.Actor{},
	}

	r := setupRouter()
	handler := NewSearchMoviesHandler(mock)
	r.GET("/search", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?q=xyz", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]storage.SearchResult
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["data"], 0)
}
