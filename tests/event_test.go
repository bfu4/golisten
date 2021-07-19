package tests

import (
	"golisten/event"
	"testing"
)

// TestEventEmission test the emission of events
func TestEventEmission(t *testing.T) {
	bus := event.CreateBus(true)
	eventName := "test-event"
	l := event.Listener{
		On: func(e *event.Event) {
			println(e.Name)
			if e.Name != eventName {
				t.Error("event names do not match!")
			}
		},
	}
	e := event.Event{
		Name: eventName,
		Id: 1,
	}
	bus.AddListener(l)
	bus.CallEvent(e)
}
