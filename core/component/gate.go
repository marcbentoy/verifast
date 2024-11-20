package component

type Gate struct {
	Ins       []Pin
	Outs      []Pin
	Operation func(*Gate)
}

func NewGate(ins []Pin, outs []Pin, op func(*Gate)) *Gate {
	return &Gate{
		Ins:       ins,
		Outs:      outs,
		Operation: op,
	}
}

func (g *Gate) Inputs() []Pin {
	return g.Ins
}

func (g *Gate) Outputs() []Pin {
	return g.Outs
}

func (g *Gate) Evaluate() {
	g.Operation(g)
}
