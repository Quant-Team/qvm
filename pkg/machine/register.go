package machine

import (
	"fmt"

	"github.com/Quant-Team/qvm/pkg/circuit"
	"github.com/Quant-Team/qvm/pkg/config"
	"github.com/Quant-Team/qvm/pkg/math/vector"
)

var _ Register = &register{}

type Register interface {
	Probability() [][]float64
	Measure() []circuit.Qubiter
}

type register struct {
	q []circuit.Qubiter
	v *vector.Vector
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

func (r *register) Vector() *vector.Vector {
	return r.v
}

func (r *register) One() {
	r.q = append(r.q, circuit.One())
	r.v = r.v.ProductVector(vector.New(0, 1))
}

func (r *register) Zero() {
	r.q = append(r.q, circuit.Zero())
	r.v = r.v.ProductVector(vector.New(1, 0))
}

func NewRegister(cfg *config.Register, state []Bit) (r *register, err error) {
	if len(state) != cfg.QubitCount {
		err = ErrInvalidCountQubit
		return
	}
	v := vector.New(1, 0)
	q := []circuit.Qubiter{circuit.Zero()}
	if state[0] == One {
		q = []circuit.Qubiter{circuit.One()}
		v = vector.New(0, 1)
	}
	r = &register{
		q: q,
		v: v,
	}

	for _, b := range state[1:] {
		switch b {
		case One:
			r.One()
		case Zero:
			r.Zero()
		}
	}
	return
}
