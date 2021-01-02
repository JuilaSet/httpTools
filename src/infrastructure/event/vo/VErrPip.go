package vo

type VErrPip struct {
	Pip chan error
}

func NewVErrPip() *VErrPip {
	return &VErrPip{
		Pip: make(chan error),
	}
}

func (e *VErrPip) Emit(err error) {
	e.Pip <- err
}

func (e *VErrPip) Bus() chan error {
	return e.Pip
}
