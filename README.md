# golisten

super simple and small event bus structure
for golang that allows emissions as go routines.


## Usage
```go
import (
	"github.com/bfu4/golisten"
)

func UseEventBus() {
	bus := golisten.CreateBus(true)

	l := golisten.Listener{
        On: func(e *golisten.Event, data ...interface{}) {
        	println(e.Name)
        	println(data[0].string())
        }
    },
    
    bus.AddListener(l)
	
	e := golisten.Event{
		Name: "event"
		Id: 1
    }
    
    bus.CallEvent(e, "hi")
}
```