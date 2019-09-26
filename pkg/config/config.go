package config

import (
	"os"
	"strconv"
)

type Register struct {
	// CountQubit in register
	// can user value COUNT_QUBIT
	CountQubit int `yaml:"count_qubit" json:"count_qubit"`
}

func FromEnv() *Register {
	countQubit, _ := strconv.Atoi(os.Getenv("COUNT_QUBIT"))
	if countQubit <= 0 {
		countQubit = 0
	}

	return &Register{
		CountQubit: countQubit,
	}
}
