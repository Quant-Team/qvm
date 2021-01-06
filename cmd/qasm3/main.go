package main

import (
	"flag"
	"fmt"
	"os"
	stdpath "path"

	"github.com/Quant-Team/qvm/pkg/machine"
)

var (
	path  string
	debug bool
)

func init() {
	flag.StringVar(&path, "file", "", "file of qasm")
	flag.BoolVar(&debug, "debug", false, "print debug token")
	flag.Parse()
}

func main() {

	dir := stdpath.Dir(path)
	register := machine.NewRegisterQasm(debug, dir)

	register.ParseProgram(path)
	w := os.Stdout
	pp := register.Probability()
	fmt.Fprintf(w, "q.Probability: %v\n", pp)
	fmt.Fprintln(w, "Measure me ^^")
	m := register.Measure()
	fmt.Fprintln(w, "q.Measure:")
	for _, q := range m {
		fmt.Fprintf(w, "%v\n", q)
	}
}
