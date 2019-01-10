package main

import (
	"fmt"
	"math"
)

func jumpSearch(arr []int, key int) int {
	step := int(math.Sqrt(float64(len(arr))))

	cur := 0
	for cur <= len(arr) {
		if arr[cur] == key {
			return cur
		}
		if arr[cur] > key {
			for i := cur; i >= cur-step; i-- {
				if arr[i] == key {
					return i
				}
			}
		}
		cur += step
	}

	return -1
}

func main() {
	sample := []int{1, 2, 3, 4, 5, 7, 8, 10, 40, 50, 54, 68, 69, 70}
	fmt.Printf("Found at: %d\n", jumpSearch(sample, 70))
}
