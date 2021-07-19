package golisten

import (
	"testing"
)

// TestEventEmission test the emission of events
func TestEventEmission(t *testing.T) {
	bus := CreateBus(true)
	eventName := "test-event"
	l := Listener{
		On: func(e *Event, data ...interface{}) {
			println(e.Name)
			// We can assert the type since we know what it is.
			println(data[0].(string))
			if e.Name != eventName {
				t.Error("event names do not match!")
			}
		},
	}
	e := Event{
		Name: eventName,
		Id: 1,
	}
	bus.AddListener(l)
	bus.CallEvent(e, "hi")
}
