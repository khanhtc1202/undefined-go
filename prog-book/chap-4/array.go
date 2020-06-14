package main

func main() {
	r := [...]int{99: -1} // array with 100 elements and the last is -1, others are 0 (init value)
	println(len(r))
	println(r[len(r) - 1])

	// arrays with the same size have the same type ([2]int for example) so comparable
	a := [2]int{1,2}
	b := [...]int{1,2}
	c := [2]int{1,3}
	println(a == b, b == c, c == a) // true false false
	// then, arrays with difference size are uncomparable :)
	// eg d := [3]int{1,2,3} a == d => panic error
}
