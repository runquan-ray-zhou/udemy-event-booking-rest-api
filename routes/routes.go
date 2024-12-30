package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// routes, part of the same package
	server.GET("/events", getEvents)          // register a handler for incoming get request
	server.GET("/events/:id", getEvent)       // get events by id
	server.POST("/events", createEvent)       // post request
	server.PUT("/events/:id", updateEvent)    // update request
	server.DELETE("/events/:id", deleteEvent) // delete request
	server.POST("/signup", signup)            // create new user route
}
