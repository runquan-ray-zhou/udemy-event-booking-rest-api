package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/runquan-ray-zhou/udemy-event-booking-rest-api/models"
)

func signup(context *gin.Context) { // sign up request handler
	var user models.User

	err := context.ShouldBindJSON(&user) // similar to Scan() function, store that request data into user, must follow structure of User

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."}) // response message if there is error
		return
	}

	err = user.Save() // store user in database

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
