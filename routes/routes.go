// routes/routes.go

package routes

import (
	"task-management-with-gin/controllers"
	"task-management-with-gin/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the main router
func SetupRouter() *gin.Engine {
	routes := gin.Default()

	// CORS middleware with custom configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization") // Add Authorization header

	routes.Use(cors.New(config))

	// Create controllers instances
	userController := controllers.NewUserController()
	boardController := controllers.NewBoardController()

	// Define your routes here
	routes.GET("/user/profile", middlewares.AuthGuard, userController.GetUserProfile)
	routes.GET("/user/refresh-token", userController.RefreshToken)
	routes.POST("/user/register", userController.Register)
	routes.POST("/user/sign-in", userController.SignIn)
	routes.PUT("/user/profile", middlewares.AuthGuard, userController.UpdateUserProfile)

	routes.POST("/board/new", middlewares.AuthGuard, boardController.NewBoard)
	routes.POST("/board/add-members", middlewares.AuthGuard, boardController.NewBoardMembers)

	return routes
}
