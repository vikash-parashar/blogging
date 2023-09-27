// controllers/auth.go
package controllers

import (
	"blogging/auth"
	"blogging/config"
	"blogging/helpers"
	"blogging/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration.
func Register(c *gin.Context) {

	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	email := c.Request.FormValue("email")
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")

	// Hash the user's password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	// Create a new user with a valid UUID for the user_id
	newUser := models.User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		Password:  string(hashedPassword),
	}
	if newUser.Email == "test@test.com" {
		newUser.Role = models.Admin
	} else {
		newUser.Role = models.General
	}

	// Check if a user with the same email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		// User with the same email already exists, return an error response
		c.JSON(http.StatusConflict, gin.H{"error": "user with this email already exists"})
		return
	}

	// Create the user if it doesn't already exist
	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

// Login handles user login.
func Login(c *gin.Context) {
	var loginRequest = models.User{
		Email:    c.Request.FormValue("email"),
		Password: c.Request.FormValue("password"),
	}
	var user models.User
	log.Println(">>>>>>>>>>>>>>>>>>>>>>.", loginRequest.Email)
	// Query the database to find the user by email
	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "please register first yourself ."})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no account found with this email ."})
		return
	}
	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password ."})
		return
	}

	// Generate a JWT token (you need to implement your JWT logic here)
	tokenString, err := auth.CreateToken(user)
	if err != nil {
		log.Println("failed to generate jwt token")
	}

	// Return the token as a response
	c.JSON(http.StatusCreated, gin.H{"message": "success", "token": tokenString})

}

// Logout handles user logout.
func Logout(c *gin.Context) {
	// Assuming you have access to the current user's token
	token := helpers.GetTokenFromRequest(c)

	// Add the token to the blacklist (you need to implement this)
	helpers.AddTokenToBlacklist(token)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
