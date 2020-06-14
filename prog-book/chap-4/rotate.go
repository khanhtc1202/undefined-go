package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateLeftNTimes(s []int, n int) {
	if n < 0 {
		panic("should bigger than 0")
	}
	if n >= len(s) {
		n = n % len(s)
	}
	// rotate left n times archives by
	// - reverse n first elements
	// - reverse the rest
	// - reverse the whole slice
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotateRightNTimes(s []int, n int) {
	if n < 0 {
		panic("should bigger than 0")
	}
	if n >= len(s) {
		n = n % len(s)
	}
	// rotate right n times archives by
	// - reverse the whole slice
	// - reverse n first elements
	// - reverse the rest
	reverse(s)
	reverse(s[:n])
	reverse(s[n:])
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotateLeftNTimes(s, 7)
	fmt.Println("Rotate left 7 time(s):", s)
	rotateRightNTimes(s, 1)
	fmt.Println("Rotate right 1 time(s):", s)
}
