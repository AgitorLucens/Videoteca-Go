package actor_test

import (
	actorcommand "be/internals/handler/actor/command"
	actorqueries "be/internals/handler/actor/queries"
	"be/internals/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockActorRepo struct {
	actors    []storage.Actor
	createErr error
	updateErr error
	deleteErr error
	getAllErr error
	created   *storage.Actor
	updated   *storage.Actor
}

func (m *mockActorRepo) GetAllActors() ([]storage.Actor, error) {
	if m.getAllErr != nil {
		return nil, m.getAllErr
	}
	return m.actors, nil
}

func (m *mockActorRepo) GetActorByID(id uint) (*storage.Actor, error) {
	for _, a := range m.actors {
		if a.ActorID == id {
			return &a, nil
		}
	}
	return nil, assert.AnError
}

func (m *mockActorRepo) CreateActor(req storage.CreateActorRequest) (*storage.Actor, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	m.created = &storage.Actor{ActorID: 1, ActorFirstName: req.FirstName, ActorLastName: req.LastName, ActorPhoto: req.Photo}
	return m.created, nil
}

func (m *mockActorRepo) UpdateActor(id uint, req storage.UpdateActorRequest) (*storage.Actor, error) {
	if m.updateErr != nil {
		return nil, m.updateErr
	}
	m.updated = &storage.Actor{ActorID: id, ActorFirstName: req.FirstName, ActorLastName: req.LastName, ActorPhoto: req.Photo}
	return m.updated, nil
}

func (m *mockActorRepo) DeleteActor(id uint) error {
	return m.deleteErr
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetActorsHandler_Success(t *testing.T) {
	photo := "data:image/png;base64,abc"
	mock := &mockActorRepo{
		actors: []storage.Actor{
			{ActorID: 1, ActorFirstName: "Keanu", ActorLastName: "Reeves", ActorPhoto: &photo},
		},
	}

	r := setupRouter()
	handler := actorqueries.NewGetActorsHandler(mock)
	r.GET("/actors", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/actors", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]storage.Actor
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["data"], 1)
	assert.Equal(t, "Keanu", response["data"][0].ActorFirstName)
}

func TestGetActorsHandler_Error(t *testing.T) {
	mock := &mockActorRepo{getAllErr: assert.AnError}

	r := setupRouter()
	handler := actorqueries.NewGetActorsHandler(mock)
	r.GET("/actors", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/actors", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateActorHandler_Success(t *testing.T) {
	mock := &mockActorRepo{}

	r := setupRouter()
	handler := actorcommand.NewCreateActorHandler(mock)
	r.POST("/actors", handler.Handle)

	body := `{"firstName":"Keanu","lastName":"Reeves","photo":"data:image/png;base64,abc"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/actors", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]*storage.Actor
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Keanu", response["data"].ActorFirstName)
	assert.Equal(t, "Reeves", response["data"].ActorLastName)
}

func TestCreateActorHandler_WithoutPhoto(t *testing.T) {
	mock := &mockActorRepo{}

	r := setupRouter()
	handler := actorcommand.NewCreateActorHandler(mock)
	r.POST("/actors", handler.Handle)

	body := `{"firstName":"Keanu","lastName":"Reeves"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/actors", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateActorHandler_InvalidBody(t *testing.T) {
	mock := &mockActorRepo{}

	r := setupRouter()
	handler := actorcommand.NewCreateActorHandler(mock)
	r.POST("/actors", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/actors", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateActorHandler_RepoError(t *testing.T) {
	mock := &mockActorRepo{createErr: assert.AnError}

	r := setupRouter()
	handler := actorcommand.NewCreateActorHandler(mock)
	r.POST("/actors", handler.Handle)

	body := `{"firstName":"Keanu","lastName":"Reeves"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/actors", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateActorHandler_Success(t *testing.T) {
	mock := &mockActorRepo{}

	r := setupRouter()
	handler := actorcommand.NewUpdateActorHandler(mock)
	r.PUT("/actors/:id", handler.Handle)

	body := `{"firstName":"Keanu","lastName":"Reeves Updated"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/actors/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]*storage.Actor
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Reeves Updated", response["data"].ActorLastName)
}

func TestUpdateActorHandler_InvalidID(t *testing.T) {
	mock := &mockActorRepo{}

	r := setupRouter()
	handler := actorcommand.NewUpdateActorHandler(mock)
	r.PUT("/actors/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/actors/abc", strings.NewReader(`{"firstName":"A","lastName":"B"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateActorHandler_RepoError(t *testing.T) {
	mock := &mockActorRepo{updateErr: assert.AnError}

	r := setupRouter()
	handler := actorcommand.NewUpdateActorHandler(mock)
	r.PUT("/actors/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/actors/1", strings.NewReader(`{"firstName":"A","lastName":"B"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteActorHandler_Success(t *testing.T) {
	mock := &mockActorRepo{}

	r := setupRouter()
	handler := actorcommand.NewDeleteActorHandler(mock)
	r.DELETE("/actors/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/actors/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteActorHandler_InvalidID(t *testing.T) {
	mock := &mockActorRepo{}

	r := setupRouter()
	handler := actorcommand.NewDeleteActorHandler(mock)
	r.DELETE("/actors/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/actors/abc", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteActorHandler_RepoError(t *testing.T) {
	mock := &mockActorRepo{deleteErr: assert.AnError}

	r := setupRouter()
	handler := actorcommand.NewDeleteActorHandler(mock)
	r.DELETE("/actors/:id", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/actors/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
