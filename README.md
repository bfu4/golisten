# golisten

super simple and small event bus structure
for golang that allows emissions as go routines.


## Usage
```go
import (
	"github.com/bfu4/golisten"
)

func UseEventBus() {
	busName := "myBus"
	eventName := "myEventName"
	bus := golisten.DemandRoutedBus(busName)

	l := golisten.ListenerFrom(eventName,
        func(e *golisten.Event) {
        	// we can use a type assertion here
        	println(e.Data[0].(string))
        	// will print hello
        },
    )
    
    bus.AddListener(l)
	e := golisten.CreateEvent(eventName, "hello")
    bus.CallEvent(e)
}
```