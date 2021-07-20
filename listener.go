package golisten

// The RegisteredListener type.
type RegisteredListener struct {
	listener
	EventName string
}

// A RegistrableListener is a listener structure that is allowed to be registered.
type RegistrableListener struct {
	// The name of the bus to register to.
	CorrespondingBus string
	// The listener function.
	On               func(e *Event, data ...interface{})
}

// Listener is a basic interface to describe
// an event listener.
//
// The listener has a singular function that takes in an event
// as a parameter. This allows the event to be manipulated
// as requested on the listener initialization.
type listener struct {
	// On listener function
	On func(e *Event, data ...interface{})
}