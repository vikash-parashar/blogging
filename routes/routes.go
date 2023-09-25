// routes.go

package routes

import (
	"blogging/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes for your application
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Public routes (no authentication required)
	publicRouter := router.Group("/public")
	{
		publicRouter.POST("/register", controllers.UserController.CreateUser)
		publicRouter.POST("/login", controllers.UserController.Login)
		// Define other public routes
	}

	// Protected routes (authentication required)
	protectedRouter := router.Group("/protected")
	protectedRouter.Use(middleware.AuthMiddleware())

	// Initialize controllers
	userController := controllers.UserController{DB: db}
	commentController := controllers.CommentController{DB: db}
	postController := controllers.PostController{DB: db}

	// Configure routes for each entity
	SetupUserRoutes(protectedRouter, &userController)
	SetupCommentRoutes(protectedRouter, &commentController)
	SetupPostRoutes(protectedRouter, &postController)

	return router
}
