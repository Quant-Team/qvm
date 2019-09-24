package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	env := os.Getenv("SECRET")
	fmt.Println(env)

	fmt.Println("Measure me ^^")

	Measure()
}

func Measure() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float32()
	fmt.Printf("%f\n", r)
}
