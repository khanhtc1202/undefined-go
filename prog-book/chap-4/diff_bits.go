package main

import (
	"crypto/sha256"
	"fmt"
)

// Note: x&=(x-1) deletes the rightmost 1-bit in x
func BitsDifference(h1, h2 *[sha256.Size]byte) int {
	n := 0
	for i := range h1 {
		for b := h1[i] ^ h2[i]; b != 0; b &= b-1 {
			n++
		}
	}
	return n
}

func main() {
	s1 := "unodostresquatro"
	s2 := "UNODOSTRESQUATRO"
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Println(BitsDifference(&h1, &h2))
}
