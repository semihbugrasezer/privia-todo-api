package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

var JwtKey = []byte("your_secret_key") // Güvenli bir anahtar kullanın

// GenerateJWT generates a JWT token with user information
func GenerateJWT(userID uint, userType string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		UserID:   userID,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
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
