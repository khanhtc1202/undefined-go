package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.RWMutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.RUnlock()
}
