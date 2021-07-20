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
	bus := golisten.DemandRoutedBus(busName)

	l := golisten.RegistrableListener{
		CorrespondingBus: busName,
        On: func(e *golisten.Event, data ...interface{}) {
        	println(e.Name)
        	// we can use a type assertion here
        	println(data[0].(string))
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