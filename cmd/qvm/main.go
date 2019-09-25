package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
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

	fmt.Fprintln(w, "I'm virtual qubit")
	fmt.Fprintln(w, "Measure me ^^")

	Measure(w)
}

func Measure(w io.Writer) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float32()
	fmt.Fprintf(w, "%f\n", r)
}
