package event

import "httpTools/src/infrastructure/event/vo"

type IBus interface {
	Emit(ev *vo.VEvent)
	Bus() chan *vo.VEvent
}

type IErrBus interface {
	Emit(err error)
	Bus() chan error
}
