package circuit

import (
	"testing"

	"github.com/Quant-Team/qvm/pkg/circuit/gates/one"
)

func TestQubitApplyHGate(t *testing.T) {

	var q Qubiter
	q = Zero()
	q = q.Apply(one.H())

	expQ := NewQubit(complex128(0.7071067811865476+0i), complex128(-0.7071067811865476+0i))

	if !q.Equal(expQ) {
		t.Fatalf("amplitudes not equals, %s, %s", q, expQ)
	}
}

func TestQubitApplyIGate(t *testing.T) {

	var q Qubiter
	q = Zero()
	q = q.Apply(one.I())

	expQ := Zero()

	if !q.Equal(expQ) {
		t.Fatalf("amplitudes not equals, %s, %s", q, expQ)
	}
}
