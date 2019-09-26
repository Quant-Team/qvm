package machine

import (
	"github.com/Quant-Team/qvm/pkg/circuit"
	"github.com/Quant-Team/qvm/pkg/config"
)

type Register struct {
	q []circuit.AQubit
}

type Bit int

var (
	One  Bit = 1
	Zero Bit = 0
)

func NewRegister(cfg config.Register, state []Bit) (r *Register) {
	r = &Register{}
	for _, b := range state {
		var q circuit.AQubit
		switch b {
		case One:
			q = circuit.One()
		default:
			q = circuit.Zero()
		}
		r.q = append(r.q, q)
	}
	return
}
