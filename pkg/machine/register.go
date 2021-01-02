package machine

import (
	"fmt"

	"github.com/Quant-Team/qvm/internal/qasm3"
	"github.com/Quant-Team/qvm/pkg/circuit"
	"github.com/Quant-Team/qvm/pkg/config"
)

var _ Register = &register{}

type Register interface {
	Probability() [][]float64
	Measure() []circuit.Qubiter
}

type register struct {
	*qasm3.Baseqasm3Listener

	debug bool
	bits  map[string]Bit
	gates map[string]struct{}
	q     []circuit.Qubiter
}

type Bit int

var (
	One  Bit = 1
	Zero Bit = 0
)

var (
	ErrInvalidCountQubit = fmt.Errorf("Invalid count qubit in config and init state")
)

func (r *register) Probability() [][]float64 {
	t := [][]float64{}
	for _, q := range r.q {
		t = append(t, q.Probability())
	}
	return t
}

func (r *register) Measure() []circuit.Qubiter {
	for i := range r.q {
		r.q[i] = r.q[i].Measure()
	}
	return r.q
}

func NewRegister(cfg *config.Register, state []Bit) (r *register, err error) {
	if len(state) != cfg.QubitCount {
		err = ErrInvalidCountQubit
		return
	}
	r = &register{q: []circuit.Qubiter{}}

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

func NewRegisterQasm(debug bool) (r *register) {
	return &register{
		debug: debug,
		q:     []circuit.Qubiter{},
		bits:  map[string]Bit{},
		gates: map[string]struct{}{},
	}
}
