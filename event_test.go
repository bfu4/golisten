package golisten

import (
	"reflect"
	"testing"
)

// TestEventEmission test the emission of events
func TestEventEmission(t *testing.T) {
	busName := "test"
	eventName := "test1"
	// Can also utilise DemandBus
	bus := DemandRoutedBus(busName)
	l := ListenerFrom(eventName,
		func(e *Event) {
			dataZeroType := reflect.TypeOf(e.Data[0])
			if dataZeroType.Name() != "string" {
				t.Errorf("got %s, expected string", dataZeroType.Name())
			}
		},
	)
	e := CreateEvent(eventName, "hi")
	bus.AddListener(l)
	if len(bus.Listeners()) == 0 {
		t.Error("the listener did not add!")
	}
	bus.CallEvent(e)
}
