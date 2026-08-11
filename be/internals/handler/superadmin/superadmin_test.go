package superadmin_test

import (
	superadmincommand "be/internals/handler/superadmin/command"
	superadminqueries "be/internals/handler/superadmin/queries"
	"be/internals/rbac"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct {
	users      []rbac.User
	createErr  error
	getAllErr  error
	created    *rbac.User
	createdReq *rbac.CreateUserRequest
}

//interface functions

func (m *mockUserRepo) GetUsersByRole(roleName string) ([]rbac.User, error) {
	if m.getAllErr != nil {
		return nil, m.getAllErr
	}
	return m.users, nil
}

func (m *mockUserRepo) CreateUser(req rbac.CreateUserRequest) (*rbac.User, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	m.createdReq = &req
	m.created = &rbac.User{ID: 1, Name: req.Name, Email: req.Email, UserName: req.UserName}
	return m.created, nil
}

func (m *mockUserRepo) GetUserById(id uint) (*rbac.User, error) {
	if m.getAllErr != nil {
		return nil, m.getAllErr
	}
	m.created = &rbac.User{ID: 1, Name: "test", Email: "test@test.com", UserName: "testuser"}
	return m.created, nil
}

func (m *mockUserRepo) SetUserPassword(id uint, password string) error {
	if m.createErr != nil {
		return m.createErr
	}
	m.created = &rbac.User{ID: id, Name: "test", Email: "test@test.com", UserName: "testuser"}
	m.created.Password = password
	return nil
}


func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAdminsHandler_Success(t *testing.T) {
	mock := &mockUserRepo{
		users: []rbac.User{
			{ID: 1, Name: "Admin One", UserName: "admin1", Email: "admin1@test.com"},
		},
	}

	r := setupRouter()
	handler := superadminqueries.NewGetAdminsHandler(mock)
	r.GET("/admins", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admins", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]rbac.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["data"], 1)
}

func TestGetAdminsHandler_Error(t *testing.T) {
	mock := &mockUserRepo{getAllErr: assert.AnError}

	r := setupRouter()
	handler := superadminqueries.NewGetAdminsHandler(mock)
	r.GET("/admins", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admins", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateAdminHandler_Success(t *testing.T) {
	mock := &mockUserRepo{}

	r := setupRouter()
	handler := superadmincommand.NewCreateAdminHandler(mock)
	r.POST("/admins", handler.Handle)

	body := `{"name":"Admin","email":"admin@test.com","username":"admin1","password":"pass123","passwordconfirm":"pass123","role":"admin"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admins", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]*rbac.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Admin", response["data"].Name)

	assert.Equal(t, "admin", mock.createdReq.Role)
}

func TestCreateAdminHandler_ForceAdminRole(t *testing.T) {
	mock := &mockUserRepo{}

	r := setupRouter()
	handler := superadmincommand.NewCreateAdminHandler(mock)
	r.POST("/admins", handler.Handle)

	body := `{"name":"Hacker","email":"h@test.com","username":"hacker","password":"pass123","passwordconfirm":"pass123","role":"superadmin"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admins", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	assert.Equal(t, "admin", mock.createdReq.Role)
}

func TestCreateAdminHandler_InvalidBody(t *testing.T) {
	mock := &mockUserRepo{}

	r := setupRouter()
	handler := superadmincommand.NewCreateAdminHandler(mock)
	r.POST("/admins", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admins", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateAdminHandler_RepoError(t *testing.T) {
	mock := &mockUserRepo{createErr: assert.AnError}

	r := setupRouter()
	handler := superadmincommand.NewCreateAdminHandler(mock)
	r.POST("/admins", handler.Handle)

	body := `{"name":"Admin","email":"a@t.com","username":"admin1","password":"pass123","passwordconfirm":"pass123","role":"admin"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admins", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
