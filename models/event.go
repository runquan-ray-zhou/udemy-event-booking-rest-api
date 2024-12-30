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
	VALUES (?, ?, ?, ?, ?)` // ? are place holder
	stmt, err := db.DB.Prepare(query) // stored in memory, easily reuse it, reusable

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
func GetAllEvents() ([]Event, error) { // call it to get available event
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // no need to prepare, use Query() to get back a bunch of rows., use Exec() to change things in table
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // reads content of the row it is processing

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil // return events slice
}

func GetEventByID(id int64) (*Event, error) { // query to fetch database for id // null value for a pointer is nil
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id) // use to get back one row

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}
