package event

import "log"

type Pip chan Event

type Emitter struct {
	Pip       Pip
	Errors    chan error
	Listeners map[string]func(Data)
}

func NewEmitter() *Emitter {
	return &Emitter{
		Pip:       make(chan Event),
		Errors:    make(chan error),
		Listeners: make(map[string]func(Data)),
	}
}

func (e *Emitter) Emit(name string, data Data) {
	e.Pip <- NewEvent(name, data)
}

func (e *Emitter) On(name string, handler func(data Data)) {
	e.Listeners[name] = handler
}

func (e *Emitter) StartListen() {
	go func() {
		for {
			select {
			case ev := <-e.Pip:
				h := e.Listeners[ev.Name]
				h(ev.Data)
			case err := <-e.Errors:
				log.Println("error : ", err)
			}
		}
	}()
}
