package component

import (
	"fmt"

	"github.com/marcbentoy/verifast/core/gate"
)

type Component struct {
	Ins   []*gate.Pin
	Outs  []*gate.Pin
	Gates []*gate.Gate
}

func NewComponent(ins, outs []*gate.Pin, gates []*gate.Gate) *Component {
	return &Component{
		Ins:   ins,
		Outs:  outs,
		Gates: gates,
	}
}

func (c *Component) Evaluate() {
	for _, g := range c.Gates {
		g.Evaluate()
	}
}

func (c *Component) String() string {
	return fmt.Sprintf("in: %+v | out: %+v", c.Ins, c.Outs)
}
