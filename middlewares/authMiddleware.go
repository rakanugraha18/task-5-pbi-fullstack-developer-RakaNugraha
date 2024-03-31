package middlewares

import (
	"errors"
	"net/http"
	"os"
	"task_5_pbi_btpns_RakaNugraha/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("API_SECRET"))

// Definisikan struktur Claims di sini
type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": helpers.UnauthorizedError().Error()})
			c.Abort()
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": helpers.UnauthorizedError().Error()})
			c.Abort()
			return
		}

		// Set both username and user_id to the context
		c.Set("username", claims.Username)
		c.Set("user_id", claims.UserID) // Assuming you have UserID in your claims struct

		c.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	// Mendapatkan nilai dari konteks dengan kunci "user_id"
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}

	// Memastikan nilai dari konteks sesuai dengan tipe yang diharapkan (uint)
	userIDUint, ok := userID.(uint)
	if !ok {
		return 0, errors.New("invalid user ID type in context")
	}

	return userIDUint, nil
}
