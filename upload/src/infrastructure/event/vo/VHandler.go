package vo

// alias
type HFunc = func(VData)

type VHandler struct {
	Name string
	H HFunc
}

func NewHandler(name string, h func(VData)) *VHandler {
	return &VHandler{Name: name, H: h}
}

type VHandlers struct {
	Hs map[*VHandler]bool
}

func NewVHandlers() *VHandlers {
	return &VHandlers{Hs: make(map[*VHandler]bool)}
}

func (hs *VHandlers) H(d VData) {
	for h := range hs.Hs {
		h.H(d)
	}
}

func (hs *VHandlers) IsEmpty() bool {
	return len(hs.Hs) != 0
}

func (hs *VHandlers) Add(handler *VHandler) {
	hs.Hs[handler] = true
}

func (hs *VHandlers) Remove(handler *VHandler) {
	delete(hs.Hs, handler)
}

