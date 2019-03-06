package quiz_test

import (
	"fmt"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil { // r is panic value given by panic on g()
			fmt.Println("Recovered in f", r) // clear panic and raise this flag
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // not go here because of panic was revoked in g()
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
