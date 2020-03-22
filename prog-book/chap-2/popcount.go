package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountDel(x uint64) int {
	k := 0
	for x!=0 {
		x, k = x&(x-1), k+1
	}
	return k
}

func main() {
	println(PopCount(8) == PopCountDel(8))
	println(PopCount(100) == PopCountDel(100))
	println(PopCount(256) == PopCountDel(256))
	println(PopCount(34) == PopCountDel(34))
	println(PopCount(300) == PopCountDel(300))
}
