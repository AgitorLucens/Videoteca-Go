package middleware

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecretOnce  sync.Once
	jwtSecretBytes []byte
)

func GetJwtSecret() []byte {
	jwtSecretOnce.Do(func() {
		jwtSecretBytes = []byte(os.Getenv("JWT_SECRET"))
	})
	return jwtSecretBytes
}

func GenerateToken(userID uint, email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   fmt.Sprintf("%d", userID),
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(GetJwtSecret())
}

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return GetJwtSecret(), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token claims"})
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"error": "Role not found in token claims"})
			return
		}

		if sub, ok := claims["sub"].(string); ok {
			c.Set("userID", sub)
		}
		if email, ok := claims["email"].(string); ok {
			c.Set("email", email)
		}

		c.Set("role", role)
		c.Next()
	}
}
