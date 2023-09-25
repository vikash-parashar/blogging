package controllers

import (
	"net/http"
	"yourapp/models" // Import your User model or adjust the import path as needed
	"yourapp/utils"  // Import your response utility functions

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// CreateUser handles user registration
func (uc *UserController) CreateUser(c *gin.Context) {
	// Parse request data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	// Hash the user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := uc.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Return a success response
	utils.SuccessResponse(c, user)
}

// Login handles user login
func (uc *UserController) Login(c *gin.Context) {
	// Parse request data
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	// Query the user by username
	var user models.User
	if err := uc.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid password")
		return
	}

	// Generate a JWT token and send it in the response
	token, err := utils.CreateToken(&user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Return the token as a success response
	utils.SuccessResponse(c, gin.H{"token": token})
}

// UpdateUser handles updating user information
func (uc *UserController) UpdateUser(c *gin.Context) {
	// Parse request data
	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	// Get the user ID from the request context (you should set it during authentication)
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Query the user by ID
	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	// Update user information
	user.Username = updateUser.Username

	// If a new password is provided, hash and update it
	if updateUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
			return
		}
		user.Password = string(hashedPassword)
	}

	// Save the updated user in the database
	if err := uc.DB.Save(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	// Return a success response
	utils.SuccessResponse(c, user)
}

// DeleteUser handles user deletion
func (uc *UserController) DeleteUser(c *gin.Context) {
	// Get the user ID from the request context (you should set it during authentication)
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Delete the user by ID
	if err := uc.DB.Delete(&models.User{}, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	// Return a success response
	utils.SuccessResponse(c, "User deleted successfully")
}

// GetAllUsers retrieves a list of all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	// Query all users
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	// Return the list of users as a success response
	utils.SuccessResponse(c, users)
}
