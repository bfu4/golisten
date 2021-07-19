package golisten

// Listener basic interface to describe
// an event listener.
//
// The listener has a singular function that takes in an event
// as a parameter. This allows the event to be manipulated
// as requested on the listener initialization.
type Listener struct {
	// On listener function
	On func(e *Event, data ...interface{})
}