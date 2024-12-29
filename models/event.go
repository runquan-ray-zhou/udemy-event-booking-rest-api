/*
storing event data
custom event type
custom event struct

*/

package models

import (
	"time"

	"github.com/runquan-ray-zhou/udemy-event-booking-rest-api/db"
)

type Event struct {
	ID          int64     // id of event
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int       // links the even to the user who created it
}

var events = []Event{} // slice of events

// methods to interact with the events
func (e Event) Save() error { // save method to save event to database
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close() // will close whenever this function ends
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// normal function
func GetAllEvents() []Event { // call it to get available event
	return events // return events slice
}
