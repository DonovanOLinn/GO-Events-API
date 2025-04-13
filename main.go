package main

import (
	"example.com/api/db"
	"example.com/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	// Default sets up an engine that configures a HTTP server for us
	server := gin.Default()

	routes.RegisterRoutes(server)

	// Starts and runs the server to listen for incoming requests
	server.Run(":8080") // localhost:8080
}

