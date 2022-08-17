package vhdl

import "strconv"

type Part struct {
	partType string
	inCons   map[string]string
	outCons  map[string]string
}

func NewPart(partType string, cons map[string]string) *Part {
	part := &Part{}
	part.partType = partType
	part.inCons = make(map[string]string, 0)
	part.outCons = make(map[string]string, 0)
	var ins []string
	var outs []string
	switch partType {
	case "Not", "Not16", "Copy16":
		ins = []string{"in"}
		outs = []string{"out"}
	case "And", "And16", "Or", "Or16", "Nand", "Nand16":
		ins = []string{"a", "b"}
		outs = []string{"out"}
	case "Decode16":
		ins = []string{"in"}
		outs = []string{}
		for j := 0; j < 16; j++ {
			outs = append(outs, strconv.Itoa(j))
		}
	default:
		panic("unsupported partType.")
	}
	for _, i := range ins {
		part.inCons[i] = cons[i]
	}
	for _, i := range outs {
		part.outCons[i] = cons[i]
	}
	return part
}

func (p *Part) Simulate(conValues map[string]int) {
	if p.partType == "Not" {
		o := p.outCons["out"]
		i := p.inCons["in"]
		conValues[o] = 1 & (^conValues[i])
	} else if p.partType == "Not16" {
		o := p.outCons["out"]
		i := p.inCons["in"]
		conValues[o] = 65535 & (^conValues[i])
	} else if p.partType == "And" || p.partType == "And16" {
		o := p.outCons["out"]
		a := p.inCons["a"]
		b := p.inCons["b"]
		conValues[o] = conValues[a] & conValues[b]
	} else if p.partType == "Or" || p.partType == "Or16" {
		o := p.outCons["out"]
		a := p.inCons["a"]
		b := p.inCons["b"]
		conValues[o] = conValues[a] | conValues[b]
	} else if p.partType == "Nand" {
		o := p.outCons["out"]
		a := p.inCons["a"]
		b := p.inCons["b"]
		conValues[o] = 1 & (^(conValues[a] & conValues[b]))
	} else if p.partType == "Nand16" {
		o := p.outCons["out"]
		a := p.inCons["a"]
		b := p.inCons["b"]
		conValues[o] = 65535 & (^(conValues[a] & conValues[b]))
	} else if p.partType == "Copy16" {
		i := p.inCons["in"]
		o := p.outCons["out"]
		if (conValues[i] & 1) == 0 {
			conValues[o] = 0
		} else {
			conValues[o] = 65535
		}
	} else if p.partType == "Decode16" {
		v := conValues[p.inCons["in"]]
		var n = 1
		for j := 0; j < 16; j++ {
			con := p.outCons[strconv.Itoa(j)]
			if n&v == 0 {
				conValues[con] = 0
			} else {
				conValues[con] = 1
			}
			n = n * 2
		}
	} else {
		panic("Unsupported partType")
	}
}
