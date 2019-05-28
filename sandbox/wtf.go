package main

import "fmt"

func mutabilityProb() {
	type S struct {
		A string
		B []string
	}
	x := S{"x-A", []string{"x-B"}}
	y := x // copy the struct
	y.A = "y-A"
	y.B[0] = "y-B"

	fmt.Println(x, y)
	// Outputs "{x-A [y-B]} {y-A [y-B]}" -- x was modified!
}

func slideCapacityProb() {
	doStuff := func(value []string) {
		fmt.Printf("value=%v\n", value)

		value2 := value[:]
		value2 = append(value2, "b")
		fmt.Printf("value=%v, value2=%v\n", value, value2)

		value2[0] = "z"
		fmt.Printf("value=%v, value2=%v\n", value, value2)
	}

	slice1 := make([]string, 1, 1) // length 1, capacity 1
	slice1[0] = "a"
	doStuff(slice1)
	// Output:
	// value=[a] -- ok
	// value=[a], value2=[a b] -- ok: value unchanged, value2 updated
	// value=[a], value2=[z b] -- ok: value unchanged, value2 updated

	slice10 := make([]string, 1, 10) // length 1, capacity 10
	slice10[0] = "a"
	doStuff(slice10)
	// Output:
	// value=[a] -- ok
	// value=[a], value2=[a b] -- ok: value unchanged, value2 updated
	// value=[z], value2=[z b] -- WTF?!? value changed???
	// => avoid by using copy(dst, src)
}

func main() {
	mutabilityProb()
	slideCapacityProb()
}
