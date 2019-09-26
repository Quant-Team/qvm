package machine

import (
	"fmt"

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

var (
	ErrInvalidCountQubit = fmt.Errorf("Invalid count qubit in config and init state")
)

func (r *Register) Probability() [][]float64 {
	t := [][]float64{}
	for _, q := range r.q {
		t = append(t, q.Probability())
	}
	return t
}

func (r *Register) Measure() []circuit.AQubit {
	for i := range r.q {
		r.q[i] = r.q[i].Measure()
	}
	return r.q
}

func NewRegister(cfg *config.Register, state []Bit) (r *Register, err error) {
	if len(state) != cfg.QubitCount {
		err = ErrInvalidCountQubit
		return
	}
	r = &Register{q: []circuit.AQubit{}}

	for _, b := range state {
		switch b {
		case One:
			r.q = append(r.q, circuit.One())
		case Zero:
			r.q = append(r.q, circuit.Zero())
		}
	}
	return
}
