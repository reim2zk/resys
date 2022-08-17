package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("vhdl seq mode start")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		panic("file name is missing")
	}
	fmt.Println("read file:", args[0])

	seq := ReadSeqScript(args[0])
	seq.Run(0)
}
