package matrix

import (
	"gorgonia.org/tensor"
)

type Matrix struct {
	*tensor.Dense
}

func (m1 Matrix) ProductMatrix(m2 Matrix) Matrix {
	if !m1.Shape().IsMatrix() || !m2.Shape().IsMatrix() {
		panic("should be matrix")
	}
	m := m1.Shape()
	p := m2.Shape()

	tmp := []*tensor.Dense{}
	for i := 0; i < m[0]; i++ {
		for j := 0; j < m[1]; j++ {
			cij, _ := m1.At(i, j)
			n, _ := tensor.Mul(m2, tensor.New(tensor.WithShape(1), tensor.WithBacking([]complex128{cij.(complex128)})))
			tmp = append(tmp, n.(*tensor.Dense))
		}
	}

	mv3 := []complex128{}
	for l := 0; l < len(tmp); l = l + m[0] {
		for j := 0; j < p[0]; j++ {
			for i := l; i < l+m[0]; i++ {
				for k := 0; k < p[1]; k++ {
					cjk, _ := tmp[i].At(j, k)
					mv3 = append(mv3, cjk.(complex128))
				}
			}
		}
	}
	m3 := tensor.New(tensor.WithShape(m[0]*p[0], m[1]*p[1]), tensor.WithBacking(mv3))

	return Matrix{m3}
}
