package models

import "github.com/runquan-ray-zhou/udemy-event-booking-rest-api/db"

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

// methods
func (u User) Save() error { // save user to database
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}
