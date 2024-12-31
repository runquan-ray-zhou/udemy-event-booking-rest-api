package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // with claims means with data
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // token will be valid for 2 hours
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // parse received token and extracts the information in token and check if it is a valid token
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // checking the token.Method type

		if !ok { // if different method to sign the token
			return nil, errors.New("Unexpected signing method") // make sure the token you received is correct
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid // check valid field

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) // get hold of data in Claims field

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string) method to extract email and userId
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
