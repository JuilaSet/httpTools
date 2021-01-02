package vo


type VEvent struct {
	Name string
	Data VData
}

func NewEvent(name string, data VData) *VEvent {
	return &VEvent{Name: name, Data: data}
}
