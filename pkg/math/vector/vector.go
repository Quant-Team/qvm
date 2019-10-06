package vector

import (
	"github.com/Quant-Team/qvm/pkg/circuit/gates"
	"gorgonia.org/tensor"
)

type Vector struct {
	*tensor.Dense
}

type Scalar struct {
	*tensor.Dense
}

func (a *Vector) ProductVector(b Vector) Vector {
	if !a.Shape().IsVector() || !b.Shape().IsVector() {
		panic("should be vectors")
	}
	ar := a.Shape()
	br := b.Shape()
	l := ar[0] * br[0]
	k := tensor.New(tensor.WithShape(l), tensor.WithBacking(make([]complex128, ar[0]*br[0])))

	p := 0
	for i := 0; i < ar[0]; i++ {
		for j := 0; j < br[0]; j++ {
			ci, _ := a.At(i)
			cj, _ := b.At(j)
			k.Set(p, ci.(complex128)*cj.(complex128))
			p = p + 1
		}
	}
	return Vector{k}
}

func (v *Vector) Add(c complex128) *Vector {
	d, _ := v.Dense.Add(tensor.New(tensor.WithShape(1), tensor.WithBacking([]complex128{c})))
	return &Vector{d}
}

func (v *Vector) Set(i int, c complex128) *Vector {
	v.Dense.Set(i, c)
	return v
}

func (v *Vector) MulScalar(s *Scalar) *Vector {
	d, err := v.Dense.MulScalar(s.Dense, false)
	if err != nil {
		panic(err)
	}
	return &Vector{d}
}

func (v *Vector) ApplyGate(s gates.Gate) *Vector {
	ij := s.Shape()

	l := []complex128{}
	for i := 0; i < ij[0]; i++ {
		tmp := complex(0, 0)
		iterator := v.Iterator()
		for j, err := iterator.Next(); err == nil; j, err = iterator.Next() {
			vAmplitude, _ := v.At(j)
			mAmplitude, _ := s.At(i, j)
			tmp = tmp + mAmplitude.(complex128)*vAmplitude.(complex128)
		}
		l = append(l, tmp)
	}

	return &Vector{tensor.New(tensor.WithShape(ij[0]), tensor.WithBacking(l))}
}

func (v *Vector) Size() int {
	size := v.Dense.Shape()
	return size[0]
}

func New(z ...complex128) *Vector {
	v := &Vector{tensor.New(tensor.WithShape(len(z)), tensor.WithBacking(z))}
	return v
}

func NewZero(n int) *Vector {
	v := []complex128{}
	for i := 0; i < n; i++ {
		v = append(v, complex128(0+0i))
	}

	return &Vector{tensor.New(tensor.WithShape(n), tensor.WithBacking(v))}

}

func NewScalar(c complex128) *Scalar {
	d := tensor.New(tensor.WithShape(1), tensor.WithBacking([]complex128{c}))
	return &Scalar{d}
}
