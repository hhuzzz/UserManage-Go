package routes

import (
	"hello/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController) {
	// HTML routes
	r.GET("/", userController.IndexPage)
	r.GET("/users/new", userController.CreatePage)
	r.GET("/users/:id/edit", userController.EditPage)

	// API routes
	api := r.Group("/api")
	{
		api.GET("/users", userController.GetAllUsers)
		api.POST("/users", userController.CreateUser)
		api.GET("/users/:id", userController.GetUserByID)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
	}

	// Form submission routes (for non-JAX submissions)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
