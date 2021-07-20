package golisten

// The modeling for the bus system tries to be similar to this java code:
//
// ```
// public static Bus<T> demandBus<T>(final Class<T> type) {
//		Bus<?> bus = buses.get(type); // buses is a Map<Class<?>, Bus<?>>
// 		if (bus != null) {
//			return (Bus<T>) bus;
//		}
//		ArrayList<Listener<T>> listeners = new ArrayList<>();
//		return new Bus(listeners, true);
// }
// ```
// However, this level of type safety does not exist in go.
// Since what we can work with are strings, we can use strings as registration
// over a static type (similar to Java's *.class).
//
// Utilising strings implements https://github.com/bfu4/golisten/pull/1 at a higher level
// then the invocation of each listener function, such as initial creation / registration.

var (
	buses = make(map[string]Bus)
)

// Bus base structure for an event bus
// that may emit events to
// registered listeners.
type Bus struct {
	registeredName string
	allowRoutines bool
	listeners     []RegisteredListener
}

// DemandBus demands a bus without goroutine capabilities.
func DemandBus(registration string) Bus {
	// Check if a bus exists. If it does, return. Otherwise, create a new bus.
	if val, exists := buses[registration]; exists {
		return val
	} else {
		return demand(registration, false)
	}
}

// DemandRoutedBus demands a bus with goroutine capabilities.
func DemandRoutedBus(registration string) Bus {
	// Check if a bus exists. If it does, return. Otherwise, create a new bus.
	if val, exists := buses[registration]; exists {
		return val
	} else {
		return demand(registration, true)
	}
}

func demand(registration string, allowRoutines bool) Bus {
	return createBus(registration, allowRoutines)
}

// createBus Create an event bus.
func createBus(registration string, allowRoutines bool) Bus {
	return Bus{
		registeredName: registration,
		allowRoutines: allowRoutines,
		listeners:     make([]RegisteredListener, 0),
	}
}

// AddListeners add listeners to the event bus.
func (bus *Bus) AddListeners(listeners ...RegistrableListener) {
	for _, v := range listeners {
		bus.AddListener(v)
	}
}

// AddListener add a listener to the event bus.
func (bus *Bus) AddListener(rl RegistrableListener) {
	// Check the registration with the registrable listener
	// If the names do not match, the should be ignored.
	if rl.CorrespondingBus != bus.registeredName {
		return
	}
	l := listener{
		On: rl.On,
	}
	registered := RegisteredListener{
		listener:  l,
		EventName: bus.registeredName,
	}
	bus.listeners = append(bus.listeners, registered)
}

// CallEvent call the specified event.
func (bus *Bus) CallEvent(e Event, data ...interface{}) {
	// Don't allow the calling of an event when there are no listeners
	if len(bus.listeners) < 1 {
		return
	}
	// If the bus allows emission as a go routine, we can emit them as a routine.
	if bus.allowRoutines {
		go doEmit(e, data, bus.listeners...)
	} else {
		// Otherwise, run the emission not as a routine.
		doEmit(e, data, bus.listeners...)
	}
}

func (bus *Bus) Listeners() []RegisteredListener {
	return bus.listeners
}

// doEmit do the actual event emission through this call.
// this is in a separate function to be used as a go routine if specified.
func doEmit(e Event, data []interface{}, listeners ...RegisteredListener) {
	for _, l := range listeners {
		l.On(&e, data...)
	}
}
