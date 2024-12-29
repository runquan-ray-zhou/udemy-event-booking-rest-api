/*
storing event data
custom event type
custom event struct

*/

package models

import "time"

type Event struct {
	ID          int // id of event
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int //links the even to the user who created it
}

var events = []Event{} // slice of events

// methods to interact with the events
func (e Event) Save() { // save method to save event to database

	events = append(events, e)

}
