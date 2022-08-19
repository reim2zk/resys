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
}

func (p *Circuit) AddPart(partType string, cons map[string]string) {
	part := NewPart(partType, cons)
	for _, con := range part.inCons {
		p.conValues[con] = -1
	}
	for _, con := range part.outCons {
		p.conValues[con] = -1
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
			// search part
			part := p.parts[i]
			var target = true
			// 1. all inputs are evaluated
			for _, con := range part.inCons {
				target = target && (p.conValues[con] != -1)
			}
			// 2. output is not evaluated
			for _, con := range part.outCons {
				target = target && (p.conValues[con] == -1)
			}
			if target {
				part.Simulate(p.conValues)
				update = true
			}
		}
		if !update {
			break
		}
	}

	// post run
	for i := range p.parts {
		p.parts[i].PostRun(p.conValues)
	}

	// check all output are evaluated
	for _, con := range outCons {
		if p.conValues[con] == -1 {
			panic(fmt.Sprintf("unevaluated output exists. con=%s", con))
		}
	}
	return
}

func (p *Circuit) ConValue(con string) int {
	return p.conValues[con]
}

func (p *Circuit) FindFirstPart(partType string) *Part {
	for _, part := range p.parts {
		if part.partType == partType {
			return part
		}
	}
	return nil
}
