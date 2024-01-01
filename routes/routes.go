// routes/routes.go

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the main router
func SetupRouter() *gin.Engine {
	routes := gin.Default()

	// Define your routes here
	routes.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return routes
}
