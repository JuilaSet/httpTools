package event

type Data interface{}

type Event struct {
	Name string
	Data Data
}

func NewEvent(name string, data Data) Event {
	return Event{Name: name, Data: data}
}
