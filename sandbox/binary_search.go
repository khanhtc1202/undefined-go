package main

import (
	"fmt"
)

func searchInt(array []int, key int) int {
	l := 0
	r := len(array) - 1
	for l <= r {
		m := (l + r) / 2
		if array[m] == key {
			return m
		}
		if array[m] < key {
			l = m + 1
			continue
		}
		if array[m] > key {
			r = m - 1
			continue
		}
	}
	return -1
}

func main() {
	sample := []int{1, 2, 3, 4, 5, 7, 8, 10, 40, 50, 54, 68, 69, 70}
	fmt.Printf("Found at: %d\n", searchInt(sample, 1))
}
