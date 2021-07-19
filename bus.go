package main

// Bus base structure for an event bus
// that may emit events to
// registered listeners.
type Bus struct {
	allowRoutines bool
	listeners     []Listener
}

// CreateBus Create an event bus.
func CreateBus(allowRoutines bool) Bus {
	return Bus{
		allowRoutines: allowRoutines,
		listeners:     make([]Listener, 0),
	}
}

// AddListeners add listeners to the event bus.
func (bus *Bus) AddListeners(listeners ...Listener) {
	for _, v := range listeners {
		bus.AddListener(v)
	}
}

// AddListener add a listener to the event bus.
func (bus *Bus) AddListener(listener Listener) {
	bus.listeners = append(bus.listeners, listener)
}

// CallEvent call the specified event.
func (bus *Bus) CallEvent(e Event) {
	// Don't allow the calling of an event when there are no listeners
	if len(bus.listeners) < 1 {
		return
	}
	// If the bus allows emission as a go routine, we can emit them as a routine.
	if bus.allowRoutines {
		go doEmit(e, bus.listeners...)
	} else {
		// Otherwise, run the emission not as a routine.
		doEmit(e, bus.listeners...)
	}
}

// doEmit do the actual event emission through this call.
// this is in a separate function to be used as a go routine if specified.
func doEmit(e Event, listeners ...Listener) {
	for _, listener := range listeners {
		listener.On(&e)
	}
}
