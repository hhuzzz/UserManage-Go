package routes

import (
	"hello/auth"
	"hello/controllers"
	"hello/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, authController *controllers.AuthController, jwtManager *auth.JWTManager) {
	// HTML routes
	r.GET("/", userController.IndexPage)
	r.GET("/login", authController.LoginPage)
	r.GET("/users/new", userController.CreatePage)
	r.GET("/users/:id/edit", userController.EditPage)

	// API routes
	api := r.Group("/api")
	{
		// Auth routes (public)
		api.POST("/auth/register", authController.Register)
		api.POST("/auth/login", authController.Login)
		api.POST("/auth/logout", authController.Logout)

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(jwtManager))
		{
			protected.GET("/auth/me", authController.GetCurrentUser)
			protected.POST("/auth/change-password", authController.ChangePassword)
			protected.GET("/users", userController.GetAllUsers)
			protected.POST("/users", userController.CreateUser)
			protected.GET("/users/:id", userController.GetUserByID)
			protected.PUT("/users/:id", userController.UpdateUser)
			protected.DELETE("/users/:id", userController.DeleteUser)
		}
	}

	// Form submission routes (for non-AJAX submissions)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
