// comment_routes.go

package routes

import (
	"yourapp/controllers"

	"github.com/gin-gonic/gin"
)

// SetupCommentRoutes configures comment-related routes
func SetupCommentRoutes(router *gin.RouterGroup, commentController *controllers.CommentController) {
	router.POST("/comments", commentController.CreateComment)
	router.GET("/comments/:id", commentController.GetCommentByID)
	// Define other comment-related routes here
}
