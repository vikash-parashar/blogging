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
		post.POST("/new", authMiddleware, controllers.CreatePost)
		post.PUT("/edit", authMiddleware, controllers.UpdatePost)
		post.DELETE("/delete", authMiddleware, controllers.DeletePost)
		post.GET("/all", authMiddleware, controllers.GetAllPosts)
	}

	// Routes for comments
	comment := r.Group("/comments")
	{
		comment.POST("/new", authMiddleware, controllers.CreateComment)
		comment.PUT("/edit", authMiddleware, controllers.UpdateComment)
		comment.DELETE("/delete", authMiddleware, controllers.DeleteComment)
		comment.GET("/all", authMiddleware, controllers.GetAllComments)
	}

	return r
}
