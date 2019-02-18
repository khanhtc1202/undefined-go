package main

import "net/http"
import "time"

var locked bool

func handler(w http.ResponseWriter, r *http.Request) {
	if locked {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	locked = true
	go time.AfterFunc(time.Millisecond, func() { locked = false })
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
