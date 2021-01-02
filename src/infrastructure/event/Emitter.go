package event

import (
	"httpTools/src/infrastructure/event/model"
	"httpTools/src/infrastructure/event/vo"
	"log"
)

type Emitter struct {
	Pip       *vo.VPip
	Errors    *vo.VErrPip
	Listeners *model.Listeners
}

func NewEmitter() *Emitter {
	e := &Emitter{
		Pip:       vo.NewVPip(),
		Errors:    vo.NewVErrPip(),
		Listeners: model.NewVListener(),
	}
	return e
}

func (e *Emitter) Emit(name string, data vo.VData) {
	e.Pip.Emit(vo.NewEvent(name, data))
}

func (e *Emitter) On(name string, h vo.HFunc) {
	e.Listeners.Register(vo.NewHandler(name, h))
}

func (e *Emitter) Remove(handler *vo.VHandler) {
	e.Listeners.UnRegister(handler)
}

func (e *Emitter) Start() {
	go func() {
		for {
			select {
			case ev := <-e.Pip.Bus():
				e.Listeners.Handle(ev)
			case err := <-e.Errors.Bus():
				log.Fatal("emitter fatal --- ", err)
			}
		}
	}()
}
