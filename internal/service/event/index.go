package event

type Event struct {
	eventRepo Repo
}

func New(eventRepo Repo) *Event {
	return &Event{
		eventRepo: eventRepo,
	}
}
