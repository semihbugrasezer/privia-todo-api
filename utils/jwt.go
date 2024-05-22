package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID   uint   `json:"userId"`
	UserType string `json:"userType"`
	jwt.RegisteredClaims
}

var JwtKey = []byte("your_secret_key") // Güvenli bir anahtar kullanın

// GenerateJWT generates a JWT token with user information
func GenerateJWT(userID uint, userType string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		UserID:   userID,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// ParseJWT parses and validates a JWT token
func ParseJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
