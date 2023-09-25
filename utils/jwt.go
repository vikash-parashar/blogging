package utils

import (
	"blogging/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Your JWT secret key (keep it secret and secure)
var jwtSecret = []byte("your-secret-key")

// CreateToken generates a JWT token for a user
func CreateToken(user *models.User) (string, error) {
	// Define the token claims
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":   time.Now().Unix(),
	}

	// Create the token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses and verifies a JWT token
func ParseToken(tokenString string) (*jwt.Token, *jwt.MapClaims, error) {
	// Parse the token with the secret key
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, nil, jwt.ErrSignatureInvalid
	}

	// Extract claims from the token
	if claims, ok := token.Claims.(*jwt.MapClaims); ok {
		return token, claims, nil
	}

	return nil, nil, jwt.ErrInvalidClaims
}
