package main

func main() {
	a, b := 1, 2
	a, b = swap(b, a)
	println(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}

