package genre_test

import (
	genrecommand "be/internals/handler/genre/command"
	genrequeries "be/internals/handler/genre/queries"
	"be/internals/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockGenreRepo struct {
	genres    []storage.Genre
	createErr error
	updateErr error
	deleteErr error
	getAllErr error
}

func (m *mockGenreRepo) GetAllGenres() ([]storage.Genre, error) {
	if m.getAllErr != nil {
		return nil, m.getAllErr
	}
	return m.genres, nil
}

func (m *mockGenreRepo) GetGenreByID(id uint) (*storage.Genre, error) {
	for _, g := range m.genres {
		if g.GenreID == id {
			return &g, nil
		}
	}
	return nil, assert.AnError
}

func (m *mockGenreRepo) CreateGenre(req storage.CreateGenreRequest) (*storage.Genre, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	return &storage.Genre{GenreID: 1, GenreName: req.Name}, nil
}

func (m *mockGenreRepo) UpdateGenre(id uint, req storage.UpdateGenreRequest) (*storage.Genre, error) {
	if m.updateErr != nil {
		return nil, m.updateErr
	}
	return &storage.Genre{GenreID: id, GenreName: req.Name}, nil
}

func (m *mockGenreRepo) DeleteGenre(id uint) error {
	return m.deleteErr
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetGenresHandler_Success(t *testing.T) {
	mock := &mockGenreRepo{
		genres: []storage.Genre{
			{GenreID: 1, GenreName: "Action"},
			{GenreID: 2, GenreName: "Drama"},
		},
	}

	r := setupRouter()
	handler := genrequeries.NewGetGenresHandler(mock)
	r.GET("/genres", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/genres", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]storage.Genre
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["data"], 2)
	assert.Equal(t, "Action", response["data"][0].GenreName)
}

func TestGetGenresHandler_Error(t *testing.T) {
	mock := &mockGenreRepo{getAllErr: assert.AnError}

	r := setupRouter()
	handler := genrequeries.NewGetGenresHandler(mock)
	r.GET("/genres", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/genres", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateGenreHandler_Success(t *testing.T) {
	mock := &mockGenreRepo{}

	r := setupRouter()
	handler := genrecommand.NewCreateGenreHandler(mock)
	r.POST("/genres", handler.Handle)

	body := `{"name":"Thriller"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/genres", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]*storage.Genre
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Thriller", response["data"].GenreName)
}

func TestCreateGenreHandler_InvalidBody(t *testing.T) {
	mock := &mockGenreRepo{}

	r := setupRouter()
	handler := genrecommand.NewCreateGenreHandler(mock)
	r.POST("/genres", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/genres", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateGenreHandler_RepoError(t *testing.T) {
	mock := &mockGenreRepo{createErr: assert.AnError}

	r := setupRouter()
	handler := genrecommand.NewCreateGenreHandler(mock)
	r.POST("/genres", handler.Handle)

	body := `{"name":"Thriller"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/genres", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateGenreHandler_Success(t *testing.T) {
	mock := &mockGenreRepo{}

	r := setupRouter()
	handler := genrecommand.NewUpdateGenreHandler(mock)
	r.PUT("/genres/:id", handler.Handle)

	body := `{"name":"Comedy"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/genres/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]*storage.Genre
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), response["data"].GenreID)
	assert.Equal(t, "Comedy", response["data"].GenreName)
}

func TestUpdateGenreHandler_InvalidID(t *testing.T) {
	mock := &mockGenreRepo{}

	r := setupRouter()
	handler := genrecommand.NewUpdateGenreHandler(mock)
	r.PUT("/genres/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/genres/abc", strings.NewReader(`{"name":"X"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateGenreHandler_RepoError(t *testing.T) {
	mock := &mockGenreRepo{updateErr: assert.AnError}

	r := setupRouter()
	handler := genrecommand.NewUpdateGenreHandler(mock)
	r.PUT("/genres/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/genres/1", strings.NewReader(`{"name":"X"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteGenreHandler_Success(t *testing.T) {
	mock := &mockGenreRepo{}

	r := setupRouter()
	handler := genrecommand.NewDeleteGenreHandler(mock)
	r.DELETE("/genres/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/genres/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteGenreHandler_InvalidID(t *testing.T) {
	mock := &mockGenreRepo{}

	r := setupRouter()
	handler := genrecommand.NewDeleteGenreHandler(mock)
	r.DELETE("/genres/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/genres/abc", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteGenreHandler_RepoError(t *testing.T) {
	mock := &mockGenreRepo{deleteErr: assert.AnError}

	r := setupRouter()
	handler := genrecommand.NewDeleteGenreHandler(mock)
	r.DELETE("/genres/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/genres/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
