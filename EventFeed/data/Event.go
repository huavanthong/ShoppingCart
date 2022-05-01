package EventFeed

type Event struct {
	SequenceNumber int
	OccuredAt      string
	Name           string
}

func NewEvent(s int, o string, n string) *Event {
	return &Event{s, o, n}
}
