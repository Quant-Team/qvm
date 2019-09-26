package machine

import (
	"fmt"
	"strings"
)

var (
	ErrInvalitSymbol = fmt.Errorf("Invalid symbol in state register, valid [0|1]")
)

func ParserRegisterState(state string) ([]Bit, error) {
	b := []Bit{}
	state = strings.TrimPrefix(state, "|")
	state = strings.TrimSuffix(state, ">")
	for _, s := range state {
		switch s {
		case '0':
			b = append(b, Zero)
		case '1':
			b = append(b, One)
		default:
			return nil, ErrInvalitSymbol
		}
	}
	return b, nil
}
