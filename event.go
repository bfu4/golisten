package golisten

// Event base structure for an event.
type Event struct {
	// Name the name of the event.
	Name string
	// Data event data
	Data []interface{}
}

// CreateEvent is a shorthand for creating an event based off of
// a name and variadic data.
func CreateEvent(name string, data ...interface{}) Event {
	return Event{
		Name: name,
		Data: data,
	}
}
