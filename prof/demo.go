package main

import (
	"bytes"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

var visitors struct {
	sync.Mutex
	n int
}

var colorRx = regexp.MustCompile(`^\w*$`)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	if !colorRx.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitors.Lock()
	visitors.n++
	num := visitors.n
	visitors.Unlock()
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()
	buf.WriteString("<h1 style='color: ")
	buf.WriteString(r.FormValue("color"))
	buf.WriteString(">Welcome!</h1>You are visitor number ")
	b := strconv.AppendInt(buf.Bytes(), int64(num), 10)
	b = append(b, '!')
	w.Write(b)
}

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
