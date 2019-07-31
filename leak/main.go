package main

import (
	"fmt"
	"log"
	"net/http"
)

type tracker struct {
	r *http.Request
}

func (*tracker) Track() {
	// Do track
}

func (*tracker) Close() {
	// Do close
}

func main() {
	reqs := make(chan *http.Request, 1000)
	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		reqs <- r
		_, _ = fmt.Fprintf(w, "Ok\n")
	})

	go trackRequests(reqs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func trackRequests(reqs <-chan *http.Request) {
	for {
		select {
		case r := <-reqs:
			tracker := &tracker{r}
			defer tracker.Close()

			tracker.Track()
		}
	}
}
