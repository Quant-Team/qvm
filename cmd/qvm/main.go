package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("SECRET")
	fmt.Fprintln(w, env)

	fmt.Fprintln(w, "Measure me ^^")

	Measure(w)
}

func Measure(w io.Writer) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float32()
	fmt.Fprintf(w, "%f\n", r)
}
