package main

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	var x, y map[string]int
	x = map[string]int{"a": 1, "b": 2}
	y = map[string]int{"a": 1, "b": 2}
	println(mapEqual(x, y))
}
