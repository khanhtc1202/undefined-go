package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello Hoang Vu Tuan")
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8000", nil)
}
