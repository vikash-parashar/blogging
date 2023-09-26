package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Define a global slice to store invalidated tokens (a simple in-memory storage)
var tokenBlacklist []string

// addToBlacklist adds a token to the blacklist
func AddTokenToBlacklist(token string) {
	// Check if the token is already in the blacklist
	for _, t := range tokenBlacklist {
		if t == token {
			return // Token is already blacklisted, no need to add it again
		}
	}

	// If the token is not in the blacklist, add it
	tokenBlacklist = append(tokenBlacklist, token)
}

// isTokenBlacklisted checks if a token is in the blacklist
func IsTokenBlacklisted(token string) bool {
	for _, blacklistedToken := range tokenBlacklist {
		if blacklistedToken == token {
			return true
		}
	}
	return false
}

// getTokenFromRequest extracts the token from the Authorization header
func GetTokenFromRequest(c *gin.Context) string {
	// Get the Authorization header value
	authHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is empty
	if authHeader == "" {
		return "" // No token found
	}

	// Expecting the header value to be in the format "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "" // Invalid format
	}

	// Return the token part
	return parts[1]
}
