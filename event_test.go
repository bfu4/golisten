package main

import (
	"testing"
)

// TestEventEmission test the emission of events
func TestEventEmission(t *testing.T) {
	bus := CreateBus(true)
	eventName := "test-event"
	l := Listener{
		On: func(e *Event) {
			println(e.Name)
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
	bus.CallEvent(e)
}
