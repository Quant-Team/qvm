package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Quant-Team/qvm/pkg/config"
	"github.com/Quant-Team/qvm/pkg/machine"
	"github.com/gorilla/mux"
)

var (
	BuildTime = ""
	GitSHA    = ""
)

var cfg *config.Register

func init() {
	cfg = config.FromEnv()
	fmt.Println("config loaded: ", *cfg)
}

func main() {
	fmt.Println("sha", GitSHA)
	fmt.Println("buildTime", BuildTime)

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, GitSHA)
	fmt.Fprintln(w, BuildTime)
	fmt.Fprintf(w, "Time: %v\n", time.Now())

	state := r.URL.Query().Get("state")
	fmt.Println("input state:", state)

	bstate, err := machine.ParserRegisterState(state)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	register, err := machine.NewRegister(cfg, bstate)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	p := register.Probability()
	fmt.Fprintf(w, "q.Probability: %v\n", p)
	fmt.Fprintln(w, "Measure me ^^")
	m := register.Measure()
	v := register.Vector()
	fmt.Fprintln(w, "q.Measure:")
	for _, q := range m {
		fmt.Fprintf(w, "%v\n", q)
	}
	fmt.Fprintf(w, "size: %d q.Vector:%s\n", v.Size(), v)
}
