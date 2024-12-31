package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/runquan-ray-zhou/udemy-event-booking-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// routes, part of the same package
	server.GET("/events", getEvents)    // register a handler for incoming get request
	server.GET("/events/:id", getEvent) // get events by id

	authenticated := server.Group("/")               // create a group of routes
	authenticated.Use(middlewares.Authenticate)      // middle ware executed before route handlers executed
	authenticated.POST("/events", createEvent)       // post request, should be protected
	authenticated.PUT("/events/:id", updateEvent)    // update request, should be protected
	authenticated.DELETE("/events/:id", deleteEvent) // delete request, should be protected

	server.POST("/signup", signup) // create new user route
	server.POST("/login", login)   // handle user login
}
