package golisten

// Event base structure for an event.
type Event struct {
	// Name the name of the event.
	Name string
	// Data event data
	Data []interface{}
}

// Data type alias for an array of anything.
type Data []interface{}

// DataOf The function DataOf shorthands the
// creation of a data object from variadic arguments.
func DataOf(data ...interface{}) Data {
	return data
}

// CreateEvent is a shorthand for creating an event based off of
// a name and variadic data.
func CreateEvent(name string, data ...interface{}) Event {
	return Event{
		Name: name,
		Data: DataOf(data),
	}
}
