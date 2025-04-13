package main

import (
	"fmt"
	"net/http"

	"example.com/api/db"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	// Default sets up an engine that configures a HTTP server for us
	server := gin.Default()

	// Sets up a handler for an incoming Get request
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	// Starts and runs the server to listen for incoming requests
	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// works similar to scan function from FMT
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later"})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
