package main

import (
	"fmt"
	"net/http"
	"os"
	"task-management-with-gin/configs"
	"task-management-with-gin/helpers"
	"task-management-with-gin/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		helpers.ErrorPanic(err)
	}

	// Init database connections
	db := configs.InitPostgresConnection()
	configs.MigrateDatabase(db)

	// Initialize the Gin router
	routes := routes.SetupRouter()

	// Start the server
	serverPort := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", serverPort)
	server := &http.Server{
		Addr:    addr,
		Handler: routes,
	}

	err := server.ListenAndServe()
	helpers.ErrorPanic(err)
}
