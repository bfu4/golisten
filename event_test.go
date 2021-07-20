package golisten

import (
	"testing"
)

// TestEventEmission test the emission of events
func TestEventEmission(t *testing.T) {
	busName := "test"
	bus := DemandRoutedBus(busName)
	l := RegistrableListener{
		CorrespondingBus: busName,
		On: func(e *Event, data ...interface{}) {
			println(e.Name)
			// We can assert the type since we know what it is.
			println(data[0].(string))
		},
	}
	e := CreateEvent(busName, "hi")
	bus.AddListener(l)
	if len(bus.Listeners()) == 0 {
		t.Error("the listener did not add!")
	}
	bus.CallEvent(e)
}

func TestNoRegistration(t *testing.T) {
	bus := DemandRoutedBus("test")
	l := RegistrableListener{
		CorrespondingBus: "not-a-test",
		On:               func(e *Event, data ...interface{}) {},
	}
	bus.AddListener(l)
	if len(bus.Listeners()) != 0 {
		t.Error("a listener registered that should not have registered")
	}
}
