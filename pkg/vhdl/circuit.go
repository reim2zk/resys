package vhdl

import (
	"fmt"
)

type Circuit struct {
	conValues map[string]int
	parts     []*Part
}

func NewCircuit() *Circuit {
	p := &Circuit{}
	p.conValues = make(map[string]int)
	p.parts = make([]*Part, 0)
	return p
}

func (p *Circuit) Print() {
	fmt.Println("Circuit")
	fmt.Println("conValues:")
	fmt.Println(p.conValues)
	// fmt.Println("parts:")
	// for i := 0; i < len(p.parts); i++ {
	// 	fmt.Println(p.parts[i].partType)
	// }
}

func (p *Circuit) AddPart(partType string, cons map[string]string) {
	part := &Part{}
	part.partType = partType
	part.inCons = make(map[string]string, 0)
	part.outCons = make(map[string]string, 0)
	if partType == "Not" {
		i := cons["in"]
		o := cons["out"]
		part.inCons["in"] = i
		part.outCons["out"] = o
		p.conValues[i] = -1
		p.conValues[o] = -1
	} else if partType == "And" {
		a := cons["a"]
		b := cons["b"]
		o := cons["out"]
		part.inCons["a"] = a
		part.inCons["b"] = b
		part.outCons["out"] = o
		p.conValues[a] = -1
		p.conValues[b] = -1
		p.conValues[o] = -1
	} else if partType == "Or" {
		a := cons["a"]
		b := cons["b"]
		o := cons["out"]
		part.inCons["a"] = a
		part.inCons["b"] = b
		part.outCons["out"] = o
		p.conValues[a] = -1
		p.conValues[b] = -1
		p.conValues[o] = -1
	} else {
		panic("unsupported partType.")
	}
	p.parts = append(p.parts, part)
}

func (p *Circuit) Run(inConValues map[string]int, outCons []string, verbose int) {
	// set connector values to unevaluated (-1)
	for con := range p.conValues {
		p.conValues[con] = -1
	}

	// set input simulation data
	for con := range inConValues {
		p.conValues[con] = inConValues[con]
	}

	// loop until evaluation stopped
	for true {
		var update = false
		if verbose > 0 {
			fmt.Println(p.conValues)
		}
		for i := range p.parts {
			// finish when all output are evaluated
			var finish = true
			for j := range outCons {
				con := outCons[j]
				finish = finish && (p.conValues[con] != -1)
			}
			if finish {
				return
			}

			// search part
			part := p.parts[i]
			var target = true
			// 1. all inputs are evaluated
			for pin := range part.inCons {
				con := part.inCons[pin]
				target = target && (p.conValues[con] != -1)
			}
			// 2. output is not evaluated
			for pin := range part.outCons {
				con := part.outCons[pin]
				target = target && (p.conValues[con] == -1)
			}
			if target {
				part.Simulate(p.conValues)
				update = true
			}
		}
		if !update {
			panic("failed to find part to change")
		}
	}
	panic("failed to simulate")
}

func (p *Circuit) ConValue(con string) int {
	return p.conValues[con]
}
