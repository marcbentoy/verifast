package main

import (
	"fmt"

	"github.com/marcbentoy/verifast/core/component"
	"github.com/marcbentoy/verifast/core/gate"
)

func main() {
	fmt.Println("verifast by marcbentoy")
	fmt.Println()
	fmt.Println()

	// pins
	// nand pins
	pAi := gate.NewPin(0)
	pBi := gate.NewPin(0)
	pCi := gate.NewPin(1)
	// and pins
	pAo := gate.NewPin(0)
	pBo := gate.NewPin(0)
	// or pins
	pCo := gate.NewPin(0)

	nandG := gate.NewGate([]*gate.Pin{pAi, pBi}, []*gate.Pin{pAo}, nandFunc)
	andG := gate.NewGate([]*gate.Pin{pAo, pCi}, []*gate.Pin{pBo}, andFunc)
	orG := gate.NewGate([]*gate.Pin{pBo, pCi}, []*gate.Pin{pCo}, orFunc)

	testComponent := component.NewComponent([]*gate.Pin{pAi, pBi, pCi}, []*gate.Pin{pCo}, []*gate.Gate{nandG, andG, orG})

	fmt.Println("component before evaluation: ", *testComponent.Outs[0])
	testComponent.Evaluate()
	fmt.Println("component after evaluation: ", *testComponent.Outs[0])

	fmt.Println("nandG: ", *nandG.Outs[0])
	fmt.Println("andG: ", *andG.Outs[0])
	fmt.Println("orG: ", *orG.Outs[0])
}

func nandFunc(g *gate.Gate) {
	if (g.Ins[0].Value == 0) && (g.Ins[1].Value == 0) {
		g.Outs[0].Value = 1
	} else {
		g.Outs[0].Value = 0
	}
}

func andFunc(g *gate.Gate) {
	if (g.Ins[0].Value == 1) && (g.Ins[1].Value == 1) {
		g.Outs[0].Value = 1
	} else {
		g.Outs[0].Value = 0
	}
}

func orFunc(g *gate.Gate) {
	if (g.Ins[0].Value == 1) || (g.Ins[1].Value == 1) {
		g.Outs[0].Value = 1
	} else {
		fmt.Println("false bai")
		g.Outs[0].Value = 0
	}
}
