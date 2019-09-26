package machine

import "testing"

func TestParserRegisterState(t *testing.T) {
	tests := []struct {
		title string
		state string
		exp   []Bit
		err   error
	}{
		{
			"two bit in zero",
			"|00>",
			[]Bit{Zero, Zero},
			nil,
		},
		{
			"two bit in One",
			"|11>",
			[]Bit{One, One},
			nil,
		},
		{
			"large register",
			"|11001100>",
			[]Bit{One, One, Zero, Zero, One, One, Zero, Zero},
			nil,
		},
		{
			"Invalid symbol in register",
			"|51001100>",
			nil,
			ErrInvalitSymbol,
		},
	}

	for _, test := range tests {
		act, err := ParserRegisterState(test.state)
		if err != test.err {
			t.Fatal(err)
		}
		if !Equal(test.exp, act) {
			t.Fatal(test.title)
		}
	}
}

func Equal(exp, act []Bit) bool {
	if len(exp) != len(act) {
		return false
	}
	for i := range exp {
		if exp[i] != act[i] {
			return false
		}
	}
	return true
}
