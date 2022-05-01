package data

/************************ Define error code ************************/
// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrEventNotFound = fmt.Errorf("Event not found")

/************************ Define structure event ************************/
type Event struct {
	SequenceNumber int
	OccuredAt      string
	Name           string
}

func NewEvent(seq int, time string, name string) *Event {

	return &Event{
		SequenceNumber: seq,
		OccuredAt:      time,
		Name:           name,
	}
}
