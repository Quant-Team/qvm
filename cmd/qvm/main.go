package main

import (
	"fmt"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"time"

	"github.com/Quant-Team/qvm/pkg/circuit"
)

var (
	BuildTime = ""
	GitSHA    = ""
)

func main() {
	fmt.Println("sha", GitSHA)
	fmt.Println("buildTime", BuildTime)

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, GitSHA)
	fmt.Fprintln(w, BuildTime)

	var q circuit.AQubit

	rand.Seed(time.Now().UnixNano())
	rnd1 := rand.Float64()
	rnd2 := rand.Float64()

	q = circuit.NewQubit(
		cmplx.Sqrt(complex(0, rnd1)),
		cmplx.Sqrt(complex(rnd2, 0)),
	)
	p := q.Probability()
	fmt.Fprintf(w, "Time: %v\n", time.Now())
	fmt.Fprintf(w, "q.Probability: %v\n", p)
	m := q.Measure()
	fmt.Fprintln(w, "Measure me ^^")
	fmt.Fprintf(w, "q.Measure: %v\n", m)
}
