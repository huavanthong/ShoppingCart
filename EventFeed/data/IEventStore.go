package data

type IEventStore interface {
	GetEvents(firstEventSequenceNumber int, lastEventSequenceNumber int) []Event
	Raise(eventName string)
}
