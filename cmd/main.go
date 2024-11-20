package main

import (
	"fmt"

	"github.com/marcbentoy/verifast/core/component"
)

func main() {
	fmt.Println("verifast by marcbentoy")

	nandIns := []component.Pin{
		*component.NewPin(0),
		*component.NewPin(1),
	}
	nandOuts := []component.Pin{
		*component.NewPin(0),
		*component.NewPin(0),
	}
	nand := component.NewGate(nandIns, nandOuts, nandFunc)
	nand.Evaluate()
}

func nandFunc(g *component.Gate) {
	if (g.Ins[0].Value == 0) && (g.Ins[1].Value == 0) {
		g.Outs[0].Value = 1
	} else {
		g.Outs[0].Value = 1
	}
	fmt.Println("this is a function inside nand with ins: ", g.Ins)
}
