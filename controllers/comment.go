// controllers/comment.go
package controllers

import (
	"blogging/config"
	"blogging/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateComment(c *gin.Context) {
	// Fetch data from the form
	// content := c.Request.FormValue("content")
	// postID := c.Request.FormValue("post_id") // Assuming you have a field for post_id in the form
	id := uuid.New()

	// Create a new Comment instance
	newComment := models.Comment{
		ID: id,
		// Content: content,
		// Post:    models.Post{ID: postID}, // Assuming PostID is the foreign key for the post
	}

	// Save the new comment to the database
	config.DB.Create(&newComment)

	// Return a success response or handle errors
	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

func UpdateComment(c *gin.Context) {
	// Fetch comment ID from URL query params
	commentID := c.DefaultQuery("id", "")

	// Fetch data from the form
	// content := c.Request.FormValue("content")

	// Find the comment by ID
	var comment models.Comment
	if err := config.DB.Where("comment_id = ?", commentID).First(&comment).Error; err != nil {
		// Handle the error (e.g., comment not found)
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Update the comment with new data
	// comment.Content = content

	// Save the updated comment to the database
	config.DB.Save(&comment)

	// Return a success response or handle errors
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func DeleteComment(c *gin.Context) {
	// Fetch comment ID from request params
	commentID := c.DefaultQuery("id", "")

	// Find the comment by ID
	var comment models.Comment
	if err := config.DB.Where("comment_id = ?", commentID).First(&comment).Error; err != nil {
		// Handle the error (e.g., comment not found)
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Delete the comment from the database
	config.DB.Delete(&comment)

	// Return a success response or handle errors
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func GetAllComments(c *gin.Context) {
	// Retrieve all comments from the database
	var comments []models.Comment
	config.DB.Find(&comments)

	// Return the list of comments as a JSON response
	c.JSON(http.StatusOK, comments)
}
