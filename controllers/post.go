// controllers/post.go
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
)

func CreatePost(c *gin.Context) {
	// Fetch data from the form
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	id := uuid.New()
	token := helpers.GetTokenFromRequest(c)
	email, err := auth.GetUserEmailFromToken(token)
	if err != nil {
		log.Println("failed to get user email from token string")
	}
	user, err := auth.GetUserByEmailID(email)
	if err != nil {
		log.Println("failed to get user from db based on email id")
	}
	// Create a new Post instance
	newPost := models.Post{
		ID:      id,
		Title:   title,
		Content: content,
		Author:  *user, // Assuming AuthorID is the foreign key for the author
	}

	// Save the new post to the database
	config.DB.Create(&newPost)

	// Return a success response or handle errors
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func UpdatePost(c *gin.Context) {

	postID := c.Query("post_id")

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
	postID := c.Param("post_id")

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
