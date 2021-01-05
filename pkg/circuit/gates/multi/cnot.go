package multi

import (
	"fmt"
	"strconv"

	"github.com/Quant-Team/qvm/pkg/circuit/gates"
	"github.com/Quant-Team/qvm/pkg/circuit/gates/one"
	"github.com/Quant-Team/qvm/pkg/math/matrix"
	"gorgonia.org/tensor"
)

type GateFunc = func() gates.Gate

func CNOT(bit int, control []int64, target int64) GateFunc {
	return func() gates.Gate {
		m := one.I()
		f := fmt.Sprintf("%s%s%s", "%0", strconv.Itoa(bit), "s")
		for i := 1; i < bit; i++ {

			g2 := one.I().Matrix
			m.Matrix = m.ProductMatrix(g2)
		}

		if !m.Shape().IsMatrix() {
			panic("expected matrix shape")
		}

		dimension := m.Shape()[0]
		b := make([]complex128, dimension*dimension)
		index := make([]int, 0)
		g := matrix.Matrix{tensor.New(tensor.WithShape(dimension, dimension), tensor.WithBacking(b))}
		for i := 0; i < dimension; i++ {
			bits := []rune(fmt.Sprintf(f, strconv.FormatInt(int64(i), 2)))

			// Apply X
			apply := true
			for j := range control {
				if bits[control[j]] == '0' {
					apply = false
					break
				}
			}

			if apply {
				if bits[target] == '0' {
					bits[target] = '1'
				} else {
					bits[target] = '0'
				}
			}

			v, err := strconv.ParseInt(string(bits), 2, 0)
			if err != nil {
				panic(fmt.Sprintf("parse int: %v", err))
			}

			index = append(index, int(v))
		}
		fmt.Println(m)
		for i, ii := range index {
			for j := 0; j < dimension; j++ {
				v, _ := m.At(ii, j)
				g.SetAt(v, i, j)
			}
		}
		fmt.Println(g)

		return gates.Gate{g}
	}
}
