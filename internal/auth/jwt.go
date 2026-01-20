package auth

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getEnv(key string) string {
	if value := os.Getenv(key); value == "" {
		return value
	}

	return ""
}

var jwtSecret = []byte(getEnv("JWT_SECRET"))

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"Role"`
	jwt.RegisteredClaims
}

func GenerateJsonWebToken(userID, email, role string) (string, error) {

	exp := time.Now().Add(72 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Mustafa",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)

}

func ValidateJsonWebToken(tokenStr string) (*Claims, error) {
	tokenStr = strings.TrimSpace(tokenStr)

	if tokenStr == "" {
		return nil, errors.New("Token is empty")
	}
}
