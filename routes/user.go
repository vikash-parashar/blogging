// user_routes.go

package routes

import (
	"blogging/controllers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures user-related routes
func SetupUserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	router.GET("/users", userController.GetAllUsers)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
