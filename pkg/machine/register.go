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

type regPos [2]int

func (r *register) NewRegPos(length int) regPos {
	return regPos([2]int{len(r.q), length})
}

func (r regPos) GetStart() int {
	return r[0]
}

func (r regPos) GetEnd() int {
	return r[1]
}

type register struct {
	*qasm3.Baseqasm3Listener

	debug  bool
	dir    string
	bits   map[string]Bit
	gates  map[string]struct{}
	qbName map[string]regPos
	q      []circuit.Qubiter
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

func NewRegisterQasm(debug bool, dir string) (r *register) {
	return &register{
		debug:  debug,
		dir:    dir,
		q:      []circuit.Qubiter{},
		qbName: map[string]regPos{},
		bits:   map[string]Bit{},
		gates:  map[string]struct{}{},
	}
}
