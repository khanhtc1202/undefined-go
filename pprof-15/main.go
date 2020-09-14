package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

//func raceCondition() {
//	var data int
//	go func() {
//		data++
//	}()
//	if data == 0 {
//		println("value is", data)
//	}
//}
//
//func main() {
//	go func() {
//		log.Println(http.ListenAndServe("localhost:6060", nil))
//	}()
//	raceCondition()
//	time.Sleep(time.Second * 300)
//}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("start")
	for {
		fib(30)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}