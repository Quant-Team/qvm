package one

import (
	"math"

	gate "github.com/Quant-Team/qvm/pkg/circuit/gates"
	m "github.com/Quant-Team/qvm/pkg/math/matrix"
	"gorgonia.org/tensor"
)

func H() gate.Gate {
	l := []complex128{1 / math.Sqrt2, 1 / math.Sqrt2, 1 / math.Sqrt2, -1 / math.Sqrt2}

	return gate.Gate{m.Matrix{tensor.New(tensor.WithShape(2, 2), tensor.WithBacking(l))}}
}
