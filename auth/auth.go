package auth

import (
	"blogging/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Your secret key for signing the token
var secretKey = []byte(os.Getenv("JWT_SECRET"))

// CustomClaims is a custom struct that includes custom claims in the token.
type CustomClaims struct {
	UserID    uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	jwt.StandardClaims
}

func CreateToken(u models.User) (string, error) {
	// Create custom claims
	claims := CustomClaims{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Role:      u.Role,
		UserID:    u.ID,
		StandardClaims: jwt.StandardClaims{
			// Set expiration time (24 hours)
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	// Parse the token with your secret key
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// Extract custom claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid custom claims")
	}

	return claims, nil
}

func GetUserFromToken(tokenString string) (models.User, error) {
	var user models.User
	claims, err := ParseToken(tokenString)
	if err != nil {
		return user, err
	}

	// Extract the email from the custom claims
	user.ID = claims.UserID
	user.Email = claims.Email
	user.FirstName = claims.FirstName
	user.LastName = claims.LastName
	user.ID = claims.UserID
	if user.Email == "" {
		return user, fmt.Errorf("email not found in token claims")
	}

	return user, nil
}
