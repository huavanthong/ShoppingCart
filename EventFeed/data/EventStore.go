package data

// Create storage for event
type Events []*Event

/*********************************************
Implement interface from IEventStore.go
*********************************************/
// GetEvents: Returns the raw list of events.
func (e *Event) GetEventsWithin(firstEventSequenceNumber int, lastEventSequenceNumber int) Event {

}

/************************ Method for Event ************************/
/************ GET ************/
// GetEvent returns a list of events exist in database
func GetEvents() Events {
	return eventList
}

// GetEventById returns a single event which matches the id from the
// database.
// If a event is not found this function returns a EventNotFound error
func GetEventById(id int) (*Event, error) {
	i := findIndexByEventID(id)
	if id == -1 {
		return nil, ErrEventNotFound
	}

	return eventList[i], nil
}

/************ POST ************/
// AddProduct addies a product to list
func AddEvent(e Event) {
	// Auto to get increment ID
	e.SequenceNumber = getNextID()

	// Append input event to database 
	eventList = append(eventList, &e)
}

// Raise: Add a event with eventName
func Raise(eventName string) {

	// Create a new event match with eventName
	e:= NewEvent(
		getNextID(),
		time.Now().UTC().String(),
		eventName
	)

	// Append a new event to database
	eventList = append(eventList, &e)

}

/************************ Internal function for Event ************************/
func getNextID() int {
	// get ID at the last product in productList
	lp := eventList[len(eventList)-1]
	return lp.ID + 1
}

// findIndex finds the index of a event in the database
// returns -1 when no event can be found
func findIndexByEventID(id int) int {

	for i, e := range eventList {
		if e.SequenceNumber == id {
			return i
		}
	}
	
	return -1
}

/************************ Storage Event ************************/
// productList is a hard coded list of products for this
// example data source
var eventList = []*Event{
	&Event{
		SequenceNumber: 1,
		OccuredAt:      time.Now().UTC().String(),
		Name:           "Hello",
	},
}
