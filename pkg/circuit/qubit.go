package circuit

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"

	"github.com/Quant-Team/qvm/pkg/circuit/gates"
	m "github.com/Quant-Team/qvm/pkg/math/matrix"
	v "github.com/Quant-Team/qvm/pkg/math/vector"
)

var _ Qubiter = Zero()
var _ Qubiter = One()

type Gate m.Matrix

// Qubiter - abstractio interface of qubit
type Qubiter interface {
	Measure() Qubiter
	Probability() []float64
	Apply(gates.Gate) Qubiter
	Equal(Qubiter) bool
}

type Qubit struct {
	vec *v.Vector
}

func (q *Qubit) Probability() []float64 {
	list := []float64{}
	iterator := q.vec.Iterator()
	for i, err := iterator.Next(); err == nil; i, err = iterator.Next() {
		amp, _ := q.vec.At(i)
		p := math.Pow(cmplx.Abs(amp.(complex128)), 2)
		list = append(list, p)
	}

	return list
}

func (q *Qubit) Measure() Qubiter {

	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()

	plist := q.Probability()
	var sum float64
	for i, p := range plist {
		if sum <= r && r < sum+p {
			q.vec = v.NewZero(q.vec.Size())
			q.vec.Set(i, 1)
			break
		}
		sum = sum + p
	}

	return q
}

func (q *Qubit) String() string {
	return fmt.Sprintf("%s", q.vec)
}

func (q *Qubit) Normalize() *Qubit {
	var sum float64
	iterator := q.vec.Iterator()
	for i, err := iterator.Next(); err == nil; i, err = iterator.Next() {
		amp, _ := q.vec.At(i)
		sum = sum + math.Pow(cmplx.Abs(amp.(complex128)), 2)
	}

	z := 1 / math.Sqrt(sum)
	q.vec = q.vec.MulScalar(v.NewScalar(complex(z, 0)))

	return q
}

func (q *Qubit) Apply(g gates.Gate) Qubiter {
	q.vec = q.vec.ApplyGate(g)
	return q
}

func (q0 *Qubit) Equal(q Qubiter) bool {
	q1, ok := q.(*Qubit)
	if !ok {
		return false
	}
	if q0.vec.Shape()[0] != q1.vec.Shape()[0] {
		return false
	}

	iterator := q0.vec.Iterator()
	for i, err := iterator.Next(); err == nil; i, err = iterator.Next() {
		left, _ := q0.vec.At(i)
		right, _ := q1.vec.At(i)
		if left.(complex128) != right.(complex128) {
			return false
		}
	}
	return true
}

func NewQubit(z ...complex128) *Qubit {
	vec := v.New(z...)
	q := &Qubit{vec}
	q.Normalize()
	return q
}

func Zero() *Qubit {
	return &Qubit{
		vec: v.New(cmplx.Sqrt(0+0i), cmplx.Sqrt(1+0i)),
	}
}

func One() *Qubit {
	return &Qubit{
		vec: v.New(cmplx.Sqrt(1+0i), cmplx.Sqrt(0+0i)),
	}
}
