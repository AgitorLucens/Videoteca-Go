package movieserie_test

import (
	movieseriecommand "be/internals/handler/movie_serie/command"
	movieseriequeries "be/internals/handler/movie_serie/queries"
	"be/internals/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockMovieSerieRepo struct {
	movies     []storage.MovieSerie
	createErr  error
	updateErr  error
	deleteErr  error
	getAllErr  error
	getByIDErr error
	created    *storage.MovieSerie
	updated    *storage.MovieSerie
}

func (m *mockMovieSerieRepo) GetAllMoviesAndSeries() ([]storage.MovieSerie, error) {
	if m.getAllErr != nil {
		return nil, m.getAllErr
	}
	return m.movies, nil
}

func (m *mockMovieSerieRepo) GetMovieSerieByID(id uint) (*storage.MovieSerie, error) {
	if m.getByIDErr != nil {
		return nil, m.getByIDErr
	}
	for _, ms := range m.movies {
		if ms.ID == id {
			return &ms, nil
		}
	}
	return nil, assert.AnError
}

func (m *mockMovieSerieRepo) CreateMovieSerie(req storage.CreateMovieSerieRequest) (*storage.MovieSerie, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	m.created = &storage.MovieSerie{ID: 1, Title: req.Title, MSType: req.MSType}
	return m.created, nil
}

func (m *mockMovieSerieRepo) UpdateMovieSerie(id uint, req storage.UpdateMovieSerieRequest) (*storage.MovieSerie, error) {
	if m.updateErr != nil {
		return nil, m.updateErr
	}
	m.updated = &storage.MovieSerie{ID: id, Title: "Updated"}
	return m.updated, nil
}

func (m *mockMovieSerieRepo) DeleteMovieSerie(id uint) error {
	return m.deleteErr
}

func (m *mockMovieSerieRepo) GetGenresByMsID(msID uint) ([]storage.Genre, error) {
	return []storage.Genre{{GenreID: 1, GenreName: "Action"}}, nil
}

func (m *mockMovieSerieRepo) GetActorsByMsID(msID uint) ([]storage.Actor, error) {
	return []storage.Actor{}, nil
}

func (m *mockMovieSerieRepo) GetEpisodesByMsID(msID uint) ([]storage.Episode, error) {
	return []storage.Episode{}, nil
}

func (m *mockMovieSerieRepo) GetRatingDataByMsID(msID uint) (storage.RatingData, error) {
	return storage.RatingData{Average: 4.5, Votes: 10, Percentage: "80"}, nil
}

func (m *mockMovieSerieRepo) ReplaceMovieSerieGenres(msID uint, genreIDs []uint) error {
	return nil
}

func (m *mockMovieSerieRepo) ReplaceMovieSerieActors(msID uint, actorIDs []uint) error {
	return nil
}

func (m *mockMovieSerieRepo) CreateComment(msID uint, userEmail, comment string) error {
	return nil
}

func (m *mockMovieSerieRepo) DeleteComment(commentID uint) error {
	return nil
}
	
func (m *mockMovieSerieRepo) GetCommentByID(commentID uint) (*storage.Comment, error) {
	return &storage.Comment{},nil
}

func (m *mockMovieSerieRepo) GetCommentsByMsIDPaginated(msID uint, page, limit int) ([]storage.Comment, int64, error) {
	return []storage.Comment{}, 0, nil
}

func (m *mockMovieSerieRepo) GetUserRating(msID uint, userEmail string) (int, error) {
	return 0, nil
}

func (m *mockMovieSerieRepo) UpsertRating(msID uint, userEmail string, rating int) error {
	return nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetMovieSeriesHandler_Success(t *testing.T) {
	mock := &mockMovieSerieRepo{
		movies: []storage.MovieSerie{
			{ID: 1, Title: "Inception", MSType: "movie"},
		},
	}

	r := setupRouter()
	handler := movieseriequeries.NewGetMovieSeriesHandler(mock)
	r.GET("/movieseries", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movieseries", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]storage.MovieSerie
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["data"], 1)
	assert.Equal(t, "Inception", response["data"][0].Title)
}

func TestGetMovieSeriesHandler_Error(t *testing.T) {
	mock := &mockMovieSerieRepo{getAllErr: assert.AnError}

	r := setupRouter()
	handler := movieseriequeries.NewGetMovieSeriesHandler(mock)
	r.GET("/movieseries", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movieseries", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetMovieSerieByIDHandler_Success(t *testing.T) {
	mock := &mockMovieSerieRepo{
		movies: []storage.MovieSerie{
			{ID: 1, Title: "Inception", MSType: "movie"},
		},
	}

	r := setupRouter()
	handler := movieseriequeries.NewGetMovieSerieByIDHandler(mock)
	r.GET("/movieseries/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movieseries/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var data storage.MovieOrSerieData
	err := json.Unmarshal(w.Body.Bytes(), &data)
	assert.NoError(t, err)
	assert.Equal(t, "Inception", data.MovieSerie.Title)
	assert.Equal(t, "Action", data.Genres)
}

func TestGetMovieSerieByIDHandler_InvalidID(t *testing.T) {
	mock := &mockMovieSerieRepo{}

	r := setupRouter()
	handler := movieseriequeries.NewGetMovieSerieByIDHandler(mock)
	r.GET("/movieseries/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movieseries/abc", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetMovieSerieByIDHandler_NotFound(t *testing.T) {
	mock := &mockMovieSerieRepo{getByIDErr: assert.AnError}

	r := setupRouter()
	handler := movieseriequeries.NewGetMovieSerieByIDHandler(mock)
	r.GET("/movieseries/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movieseries/99", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateMovieSerieHandler_Success(t *testing.T) {
	mock := &mockMovieSerieRepo{}

	r := setupRouter()
	handler := movieseriecommand.NewCreateMovieSerieHandler(mock)
	r.POST("/movieseries", handler.Handle)

	body := `{"title":"Inception","synopsis":"A dream","releaseYear":2010,"classificationMS":"PG-13","director":"Nolan","cover":"base64data","msType":"movie"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movieseries", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateMovieSerieHandler_InvalidBody(t *testing.T) {
	mock := &mockMovieSerieRepo{}

	r := setupRouter()
	handler := movieseriecommand.NewCreateMovieSerieHandler(mock)
	r.POST("/movieseries", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movieseries", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateMovieSerieHandler_RepoError(t *testing.T) {
	mock := &mockMovieSerieRepo{createErr: assert.AnError}

	r := setupRouter()
	handler := movieseriecommand.NewCreateMovieSerieHandler(mock)
	r.POST("/movieseries", handler.Handle)

	body := `{"title":"X","synopsis":"Y","releaseYear":2020,"classificationMS":"R","director":"D","cover":"C","msType":"movie"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movieseries", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateMovieSerieHandler_Success(t *testing.T) {
	mock := &mockMovieSerieRepo{}

	r := setupRouter()
	handler := movieseriecommand.NewUpdateMovieSerieHandler(mock)
	r.PUT("/movieseries/:id", handler.Handle)

	body := `{"title":"Updated Title"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/movieseries/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateMovieSerieHandler_InvalidID(t *testing.T) {
	mock := &mockMovieSerieRepo{}

	r := setupRouter()
	handler := movieseriecommand.NewUpdateMovieSerieHandler(mock)
	r.PUT("/movieseries/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/movieseries/abc", strings.NewReader(`{"title":"X"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteMovieSerieHandler_Success(t *testing.T) {
	mock := &mockMovieSerieRepo{}

	r := setupRouter()
	handler := movieseriecommand.NewDeleteMovieSerieHandler(mock)
	r.DELETE("/movieseries/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/movieseries/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteMovieSerieHandler_InvalidID(t *testing.T) {
	mock := &mockMovieSerieRepo{}

	r := setupRouter()
	handler := movieseriecommand.NewDeleteMovieSerieHandler(mock)
	r.DELETE("/movieseries/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/movieseries/abc", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteMovieSerieHandler_RepoError(t *testing.T) {
	mock := &mockMovieSerieRepo{deleteErr: assert.AnError}

	r := setupRouter()
	handler := movieseriecommand.NewDeleteMovieSerieHandler(mock)
	r.DELETE("/movieseries/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/movieseries/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
