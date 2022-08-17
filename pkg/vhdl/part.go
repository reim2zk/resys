package vhdl

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
	if partType == "Not" {
		i := cons["in"]
		o := cons["out"]
		part.inCons["in"] = i
		part.outCons["out"] = o
	} else if partType == "And" {
		a := cons["a"]
		b := cons["b"]
		o := cons["out"]
		part.inCons["a"] = a
		part.inCons["b"] = b
		part.outCons["out"] = o
	} else if partType == "Or" {
		a := cons["a"]
		b := cons["b"]
		o := cons["out"]
		part.inCons["a"] = a
		part.inCons["b"] = b
		part.outCons["out"] = o
	} else {
		panic("unsupported partType.")
	}
	return part
}

func (p *Part) Simulate(conValues map[string]int) {
	if p.partType == "Not" {
		o := p.outCons["out"]
		i := p.inCons["in"]
		conValues[o] = 1 - conValues[i]
	} else if p.partType == "And" {
		o := p.outCons["out"]
		a := p.inCons["a"]
		b := p.inCons["b"]
		conValues[o] = conValues[a] * conValues[b]
	} else if p.partType == "Or" {
		o := p.outCons["out"]
		a := p.inCons["a"]
		b := p.inCons["b"]
		v := conValues[a] + conValues[b]
		if v > 0 {
			conValues[o] = 1
		} else {
			conValues[o] = 0
		}
	} else {
		panic("Unsupported partType")
	}
}
