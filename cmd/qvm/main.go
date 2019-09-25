package main

import (
	"fmt"
	"log"
	"net/http"

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

	q = circuit.Zero()
	p := q.Probability()
	fmt.Fprintf(w, "Time: %v", .time.Now())
	fmt.Fprintf(w, "q.Probability: %v", p)
	m := m.Measure()
	fmt.Fprintln(w, "Measure me ^^")
	fmt.Fprintln(w, "q.Measure: %v", q)
}
