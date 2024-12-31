package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// routes, part of the same package
	server.GET("/events", getEvents)          // register a handler for incoming get request
	server.GET("/events/:id", getEvent)       // get events by id
	server.POST("/events", createEvent)       // post request, should be protected
	server.PUT("/events/:id", updateEvent)    // update request, should be protected
	server.DELETE("/events/:id", deleteEvent) // delete request, should be protected
	server.POST("/signup", signup)            // create new user route
	server.POST("/login", login)              // handle user login
}
