package multi

import (
	"fmt"
	"testing"
)

func TestCNOT(t *testing.T) {
	g := CNOT(3, []int64{1}, 2)()
	fmt.Println(g.Shape())
	fmt.Println(g)
}
