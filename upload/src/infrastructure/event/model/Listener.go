package model

import "httpTools/src/infrastructure/event/vo"

type Listeners struct {
	Ls map[string]*vo.VHandlers
}

func NewVListener() *Listeners {
	return &Listeners{
		Ls: make(map[string]*vo.VHandlers),
	}
}

func (l *Listeners) Register(handler *vo.VHandler) {
	if _, ok := l.Ls[handler.Name]; !ok {
		l.Ls[handler.Name] = vo.NewVHandlers()
	}
	l.Ls[handler.Name].Add(handler)
}

func (l *Listeners) UnRegister(handler *vo.VHandler) {
	l.Ls[handler.Name].Remove(handler)
	if len(l.Ls[handler.Name].Hs) == 0 {
		delete(l.Ls, handler.Name)
	}
}

func (l *Listeners) Handle(event *vo.VEvent) {
	l.Ls[event.Name].H(event.Data)
}
