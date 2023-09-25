package middleware

import (
	"blogging/utils"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request headers
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.RespondWithError(c, 401, "Unauthorized: Missing token")
			c.Abort()
			return
		}

		// Parse and verify the JWT token
		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || !token.Valid {
			utils.RespondWithError(c, 401, "Unauthorized: Invalid token")
			c.Abort()
			return
		}

		// Extract user information from claims and set it in the context
		userID := uint(claims["user_id"].(float64)) // Assuming you store the user ID in the token claims
		c.Set("user_id", userID)

		c.Next()
	}
}
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if an error occurred during request processing
		if len(c.Errors) > 0 {
			// Log the error (you can customize this part)
			for _, err := range c.Errors {
				log.Error(err.Err)
			}

			// Respond with an error message
			utils.RespondWithError(c, 500, "Internal Server Error")
		}
	}
}
