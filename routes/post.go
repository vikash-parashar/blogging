// post_routes.go

package routes

import (
	"blogging/controllers"

	"github.com/gin-gonic/gin"
)

// SetupPostRoutes configures post-related routes
func SetupPostRoutes(router *gin.RouterGroup, postController *controllers.PostController) {
	router.POST("/posts", postController.CreatePost)
	router.GET("/posts/:id", postController.GetPostByID)
	// Define other post-related routes here
}
