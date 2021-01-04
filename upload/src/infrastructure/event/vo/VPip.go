package vo

type VPip struct {
	Pip chan *VEvent
}

func NewVPip() *VPip {
	return &VPip{make(chan *VEvent)}
}

func (e *VPip) Emit(ev *VEvent) {
	e.Pip <- ev
}

func (e *VPip) Bus() chan *VEvent {
	return e.Pip
}
