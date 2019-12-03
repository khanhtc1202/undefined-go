package main

import "fmt"

type Sample struct {
	val int
}

func swap(a, b Sample) (Sample, Sample) {
	return b, a
}

func swap_p(a, b *Sample) (*Sample, *Sample) {
	return b, a
}

func swap_px(a, b *Sample) {
	a, b = b, a
}

func main() {
	s1 := Sample{1}
	s2 := Sample{2}
	fmt.Println(s1, s2)
	s1, s2 = swap(s2, s1)
	fmt.Println(s1, s2)
	s3, s4 := swap_p(&s2, &s1)
	fmt.Println(s3, s4)
	fmt.Println(s1, s2)
	swap_px(&s1, &s2)
	fmt.Println(s1, s2)
}
