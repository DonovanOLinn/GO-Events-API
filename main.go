package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default sets up an engine that configures a HTTP server for us
	server := gin.Default()

	// Sets up a handler for an incoming Get request
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	// Starts and runs the server to listen for incoming requests
	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}