package golisten

import (
	"reflect"
	"testing"
)

// TestEventEmission test the emission of events
func TestEventEmission(t *testing.T) {
	busName := "test"
	// Can also utilise DemandBus
	bus := DemandRoutedBus(busName)
	l := ListenerFrom(busName,
		func(e *Event) {
			dataZeroType := reflect.TypeOf(e.Data[0])
			if dataZeroType.Name() != "string" {
				t.Errorf("got %s, expected string", dataZeroType.Name())
			}
		},
	)
	e := CreateEvent(busName, "hi")
	bus.AddListener(l)
	if len(bus.Listeners()) == 0 {
		t.Error("the listener did not add!")
	}
	bus.CallEvent(e)
}

func TestNoRegistration(t *testing.T) {
	bus := DemandRoutedBus("test")
	l := ListenerFrom("not-a-test", func(e *Event) {})
	bus.AddListener(l)
	if len(bus.Listeners()) != 0 {
		t.Error("a listener registered that should not have registered")
	}
}
