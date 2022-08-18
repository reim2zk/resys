package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"resys/pkg/vhdl"
)

type Seq struct {
	circuit       *vhdl.Circuit
	outputCons    []string
	conValuesList [](map[string]int)
}

type SeqScript struct {
	Load   string
	Output []string
	Eval   [](map[string]int)
}

type PartScript struct {
	Name string
	Con  map[string]string
}

type VhdlScript struct {
	Name  string
	In    []string
	Out   []string
	Parts []PartScript
}

func ReadSeqScript(fileName string) *Seq {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var script SeqScript
	json.Unmarshal(raw, &script)

	var seq Seq
	seq.outputCons = script.Output
	seq.conValuesList = script.Eval

	var fp string
	if script.Load == "." {
		fp = fileName
	} else {
		fp = filepath.Dir(fileName) + "/" + script.Load
	}
	seq.circuit = ReadCircuit(fp)
	return &seq
}

func ReadCircuit(fileName string) *vhdl.Circuit {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("failed to read file. fileName=%s", fileName))
	}
	var vhdlScript VhdlScript
	json.Unmarshal(raw, &vhdlScript)

	p := vhdl.NewCircuit()
	if len(vhdlScript.Parts) == 0 {
		panic("parts not found")
	}
	for _, part := range vhdlScript.Parts {
		p.AddPart(part.Name, part.Con)
	}
	return p
}

func (p *Seq) Run(verbose int) {
	fmt.Print("| ")
	for _, con := range p.outputCons {
		fmt.Print(con, " | ")
	}
	fmt.Println("")
	for _, conValues := range p.conValuesList {
		p.circuit.Run(conValues, p.outputCons, verbose)
		fmt.Print("| ")
		for _, con := range p.outputCons {
			// 65535: 1111 1111 1111 1111
			// 32768: 1000 0000 0000 0000
			var x = p.circuit.ConValue(con)
			if x&32768 > 0 {
				x = -(65535 & ^x + 1)
			}
			fmt.Print(x, " | ")
		}
		fmt.Println("")
	}
}
