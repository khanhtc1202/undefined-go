package main

import "fmt"

func doStubWithFmt() {
	x := 12
	fmt.Println(x)
}

func doStub() {
	x := 12
	println(x)
}

func main() {
	doStub()
	doStubWithFmt()
}
