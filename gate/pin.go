package gate

type Pin struct {
	Value int
}

func NewPin(val int) *Pin {
	return &Pin{
		Value: val,
	}
}

func (p *Pin) Update(newVal int) {
	p.Value = newVal
}
