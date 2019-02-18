package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/khanhtc1202/undefined-go/shouldbe/rng"
)

var delay = 100 * time.Millisecond
var N = 100
var r *rng.PoissonGenerator

type delayFunc func() float32

func constant() float32 {
	return 1
}
func uniform() float32 {
	return rand.Float32()
}
func poisson() float32 {
	return float32(r.Poisson(1))
}

func init() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	r = rng.NewPoissonGenerator(seed)
	go func() {
		for {
			fmt.Printf(" ")
			time.Sleep(delay)
		}
	}()
}

func makeRequest() {
	url := "http://localhost:5000/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Printf(".")
	} else {
		fmt.Printf("🔥")
	}
}

func Loop(name string, fn delayFunc) {
	fmt.Println(name)
	for i := 0; i < N; i++ {
		v := fn() * float32(delay)
		time.Sleep(time.Duration(v))
		makeRequest()
	}
	fmt.Println()
}

func main() {
	Loop("Constant", constant)
	Loop("Uniform", uniform)
	Loop("Poisson", poisson)
}
