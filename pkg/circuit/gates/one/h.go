package one

import (
	"math"

	gate "github.com/Quant-Team/qvm/pkg/circuit/gates"
	m "github.com/Quant-Team/qvm/pkg/math/matrix"
	"gorgonia.org/tensor"
)

// Hadamard gate
func H() gate.Gate {
	l := []complex128{1 / math.Sqrt2, 1 / math.Sqrt2, 1 / math.Sqrt2, -1 / math.Sqrt2}
	return gate.Gate{makeMatrix(l)}
}

// Identidy gate
func I() gate.Gate {
	l := []complex128{1, 0, 0, 1}
	return gate.Gate{makeMatrix(l)}
}

// X is Pauli x-axis gate
func X() gate.Gate {
	l := []complex128{0, 1, 1, 0}
	return gate.Gate{makeMatrix(l)}
}

// Y is Pauli y-axis gate
func Y() gate.Gate {
	l := []complex128{0, 0 - 1i, 0 + 1i, 0}
	return gate.Gate{makeMatrix(l)}
}

// Z is Pauli z-axis gate
func Z() gate.Gate {
	l := []complex128{1, 0, 0, -1}
	return gate.Gate{makeMatrix(l)}
}

func makeMatrix(l []complex128) m.Matrix {
	return m.Matrix{tensor.New(tensor.WithShape(2, 2), tensor.WithBacking(l))}
}
