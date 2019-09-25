package circuit

import (
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var _ AQubit = Zero()
var _ AQubit = One()

// AQubit - abstractio interface of qubit
type AQubit interface {
	Measure() AQubit
	Probability() []float64
}

type Vector []complex128

func New(z ...complex128) Vector {
	v := Vector{}
	for _, zi := range z {
		v = append(v, zi)
	}

	return v
}

func NewZero(n int) Vector {
	v := Vector{}
	for i := 0; i < n; i++ {
		v = append(v, 0)
	}

	return v
}

type Qubit struct {
	v Vector
}

func (q *Qubit) Probability() []float64 {
	list := []float64{}
	for _, amp := range q.v {
		p := math.Pow(cmplx.Abs(amp), 2)
		list = append(list, p)
	}

	return list
}

func (q *Qubit) Measure() AQubit {

	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()

	plist := q.Probability()
	var sum float64
	for i, p := range plist {
		if sum <= r && r < sum+p {
			q.v = NewZero(len(q.v))
			q.v[i] = 1
			break
		}
		sum = sum + p
	}

	return q
}

func Zero() *Qubit {
	return &Qubit{
		v: Vector{cmplx.Sqrt(0 + 0i), cmplx.Sqrt(1 + 1i)},
	}
}

func One() *Qubit {
	return &Qubit{
		v: Vector{cmplx.Sqrt(1 + 1i), cmplx.Sqrt(0 + 0i)},
	}
}
