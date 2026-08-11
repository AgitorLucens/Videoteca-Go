package auserauthentication

import (
	"be/internals/middleware"
	"be/internals/rbac"
	"be/util/crypt"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	r *rbac.Repository
}

func NewLoginHandler(r *rbac.Repository) *LoginHandler {
	return &LoginHandler{
		r: r,
	}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	var req rbac.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.r.GetUserByUserName(req.UserName)
	if err != nil {
		middleware.RecordLoginAttempt(false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	ok, _ := crypt.CheckPassword(req.Password, u.Password)
	if !ok {
		middleware.RecordLoginAttempt(false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString, err := middleware.GenerateToken(u.ID, u.Email, u.Roles[0].RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	middleware.RecordLoginAttempt(true)

	device := c.Request.UserAgent()
	middleware.TrackConnectedDevice(
		fmt.Sprintf("%d", u.ID),
		u.Roles[0].RoleName,
		device,
		time.Now().Add(time.Hour),
	)

	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"message": "Login successful",
	})
}
