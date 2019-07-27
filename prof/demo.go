package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var visitors struct{
	sync.Mutex
	n int
}

var colorRx = regexp.MustCompile(`^\w*$`)

func handleHi(w http.ResponseWriter, r *http.Request) {
	if !colorRx.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitors.Lock()
	visitors.n++
	visitorsNum := visitors.n
	visitors.Unlock()
	fmt.Fprintf(w, "<h1 style='color: %s'>Welcome!</h1>You are visitor number %d!",
		r.FormValue("color"), visitorsNum)
}

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
