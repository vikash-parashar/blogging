// controllers/post.go
package controllers

import (
	"blogging/auth"
	"blogging/config"
	"blogging/helpers"
	"blogging/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePost(c *gin.Context) {
	// Fetch data from the JSON request body
	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Get the user details from the token
	token := helpers.GetTokenFromRequest(c)
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Create a new Post instance and associate it with the user
	id := uuid.New()
	newPost := models.Post{
		ID:      id,
		Title:   post.Title,
		Content: post.Content,
		Author:  user, // Associate the post with the user
	}
	fmt.Println(newPost)
	// Save the new post to the database
	config.DB.Create(&newPost)

	// Return a success response or handle errors
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func UpdatePost(c *gin.Context) {

	postID := c.DefaultQuery("id", "")

	// Fetch data from the form
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")

	// Find the post by ID
	var post models.Post
	if err := config.DB.Where("post_id = ?", postID).First(&post).Error; err != nil {
		// Handle the error (e.g., post not found)
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Update the post with new data
	post.Title = title
	post.Content = content

	// Save the updated post to the database
	config.DB.Save(&post)

	// Return a success response or handle errors
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	// Fetch post ID from request params
	postID := c.DefaultQuery("id", "")

	// Find the post by ID
	var post models.Post
	if err := config.DB.Where("post_id = ?", postID).First(&post).Error; err != nil {
		// Handle the error (e.g., post not found)
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Delete the post from the database
	config.DB.Delete(&post)

	// Return a success response or handle errors
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func GetAllPosts(c *gin.Context) {
	// Retrieve all posts from the database
	var posts []models.Post
	config.DB.Find(&posts)

	// Return the list of posts as a JSON response
	c.JSON(http.StatusOK, posts)
}
