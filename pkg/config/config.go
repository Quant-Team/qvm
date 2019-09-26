package config

import (
	"os"
	"strconv"
)

type Register struct {
	// CountQubit in register
	// can user value COUNT_QUBIT
	QubitCount int `yaml:"qubit_count" json:"qubit_count"`
}

func FromEnv() *Register {
	qubitCount, _ := strconv.Atoi(os.Getenv("QUBIT_COUNT"))
	if qubitCount <= 0 {
		qubitCount = 0
	}

	return &Register{
		QubitCount: qubitCount,
	}
}
