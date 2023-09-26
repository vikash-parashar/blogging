// routes/routes.go
package routes

import (
	"blogging/controllers"
	"blogging/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize JWT middleware
	authMiddleware := middleware.AuthMiddleware()

	// Routes for authentication
	page := r.Group("")
	{
		page.GET("/", controllers.RenderHomepage)
		page.GET("/login", controllers.RenderLoginPage)
		page.GET("/register", controllers.RenderRegisterPage)
		page.POST("/register", controllers.Register)
		page.POST("/login", controllers.Login)
		page.POST("/logout", controllers.Logout)
	}

	// Routes for posts
	post := r.Group("/posts")
	{
		post.POST("/", authMiddleware, controllers.CreatePost)
		post.PUT("/:id", authMiddleware, controllers.UpdatePost)
		post.DELETE("/:id", authMiddleware, controllers.DeletePost)
		post.GET("/", controllers.GetAllPosts)
	}

	// Routes for comments
	comment := r.Group("/comments")
	{
		comment.POST("/", authMiddleware, controllers.CreateComment)
		comment.PUT("/:id", authMiddleware, controllers.UpdateComment)
		comment.DELETE("/:id", authMiddleware, controllers.DeleteComment)
		comment.GET("/", controllers.GetAllComments)
	}

	return r
}
