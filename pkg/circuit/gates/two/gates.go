package two

import (
	gate "github.com/Quant-Team/qvm/pkg/circuit/gates"
	m "github.com/Quant-Team/qvm/pkg/math/matrix"
	"gorgonia.org/tensor"
)

const (
	dimension = 4
)

// CNOT is two qubits control not gate
// like on XOR in classical logic
func CNOT() gate.Gate {
	l := []complex128{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 0, 1,
		0, 0, 1, 0,
	}
	return gate.Gate{makeMatrix(l)}
}

func makeMatrix(l []complex128) m.Matrix {
	return m.Matrix{tensor.New(tensor.WithShape(dimension, dimension), tensor.WithBacking(l))}
}
