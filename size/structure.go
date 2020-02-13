package main

import "unsafe"

type A struct {
	foo bool      // size 1 byte
	num_1 float64 // size 8 bytes
	num_2 int32   // size 4 bytes
} // expected size 13 bytes

type B struct {
	num_1 float64
	foo bool
	num_2 int32
}

func main() {
	a := A{}
	println("size of a = ", unsafe.Sizeof(a)) // output: 24 | go runtime give memory by block and it's size 8 bytes for each
	var b B
	println("size of b = ", unsafe.Sizeof(b)) // output: 16 | go runtime will fill val to data block sequently till it full :)
}

