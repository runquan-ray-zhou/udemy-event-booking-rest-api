package main

import (
	"github.com/gin-gonic/gin"
	"github.com/runquan-ray-zhou/udemy-event-booking-rest-api/db"
	"github.com/runquan-ray-zhou/udemy-event-booking-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default() // configures http server that comes with middleware, a server pointer

	routes.RegisterRoutes(server)

	server.Run("127.0.0.1:8080") // start listening to incoming request when main func is executed, currently on local host 8080

}
