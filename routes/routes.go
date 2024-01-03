// routes/routes.go

package routes

import (
	"net/http"
	"task-management-with-gin/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the main router
func SetupRouter() *gin.Engine {
	routes := gin.Default()

	// Create controllers instances
	userController := controllers.NewUserController()

	// Define your routes here
	routes.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.POST("/user/register", userController.Register)
	routes.POST("/user/sign-in", userController.SignIn)

	return routes
}
