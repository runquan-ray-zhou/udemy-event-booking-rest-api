package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/runquan-ray-zhou/udemy-event-booking-rest-api/models"
)

func main() {
	server := gin.Default() // configures http server that comes with middleware, a server pointer

	server.GET("/events", getEvents)    // register a handler for incoming get request
	server.POST("/events", createEvent) // post request

	server.Run("127.0.0.1:8080") // start listening to incoming request when main func is executed, currently on local host 8080

}

// send back a response
func getEvents(context *gin.Context) { // gin will pass a context parament to the function, if you set this handler as this endpoint.
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events) // send back a response in JSON format, pass back a http status code and data
}

func createEvent(context *gin.Context) { // extract of data from request
	var event models.Event
	err := context.ShouldBindJSON(&event) // similar to Scan() function, store that request data into event, must follow structure of Event

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."}) // response message if there is error
		return
	}

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
	// 	return
	// }

	event.ID = 1
	event.UserID = 1

	event.Save()                                                                         //
	context.JSON(http.StatusCreated, gin.H{"message": "Event created:", "event": event}) // response message when event is created successful
}
