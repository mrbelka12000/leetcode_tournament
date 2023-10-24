package models

type Index struct {
}

// EventPage оторбражается на странице /event/{id}
type EventPage struct {
	Event Event
	Usrs  []Usr
}
