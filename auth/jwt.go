package auth

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
	"os"
	"fmt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken generates a JWT token for the given user ID
func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &jwt.StandardClaims{
		Issuer:    userID,
		ExpiresAt: expirationTime.Unix(),
	}

	// Debugging: Print expiration time
	fmt.Println("Token expires at:", expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken validates the token and returns the claims
func ValidateToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Debugging: Print the error if there's one
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

// Middleware to protect routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Store the user ID in the context to be used by other handlers
		c.Set("userID", claims.Issuer)
		c.Next()
	}
}
