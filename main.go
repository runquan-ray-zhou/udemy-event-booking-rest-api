package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // configures http server that comes with middleware, a server pointer

	server.GET("/events", getEvents) // register a handler for incoming get request

	server.Run(":8080") // start listening to incoming request when main func is executed, currently on local host 8080

}

// send back a response
func getEvents(context *gin.Context) { // gin will pass a context parament to the function, if you set this handler as this endpoint.
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) // send back a response in JSON format, pass back a http status code and data
}
