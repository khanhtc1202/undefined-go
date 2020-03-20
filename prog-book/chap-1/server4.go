package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler4)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler4(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			cycles, err := strconv.Atoi(v[0])
			if err != nil {
				log.Print(err)
			}
			lissajous(w, cycles)
		}
	}
	fmt.Fprintf(w, "Missing parametter")
}
