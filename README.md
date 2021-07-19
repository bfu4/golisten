# golisten

super simple and small event bus structure
for golang that allows emissions as go routines.


## Usage
```go
import (
	"github.com/bfu4/golisten"
)

func UseEventBus() {
	bus := event.CreateBus(true)

	l := Listener{
        On: func(e *Event, data ...interface{}) {
        	println(e.Name)
        	println(data[0].string())
        }
    },
    
    bus.AddListener(l)
	
	e := event.Event{
		Name: "event"
		Id: 1
    }
    
    bus.CallEvent(e, "hi")
}
```