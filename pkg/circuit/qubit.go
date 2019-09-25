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

func (v0 Vector) Mul(z complex128) Vector {
	v2 := Vector{}
	for i := range v0 {
		v2 = append(v2, z*v0[i])
	}

	return v2
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

func (q *Qubit) Normalize() *Qubit {
	var sum float64
	for _, amp := range q.v {
		sum = sum + math.Pow(cmplx.Abs(amp), 2)
	}

	z := 1 / math.Sqrt(sum)
	q.v = q.v.Mul(complex(z, 0))

	return q
}

func NewQubit(z ...complex128) *Qubit {
	v := Vector{}
	for _, zi := range z {
		v = append(v, zi)
	}
	q := &Qubit{v}
	q.Normalize()
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
