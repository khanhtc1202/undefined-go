package main

import "fmt"

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2;; i++ {
			ch <- i
		}
	}()
	return ch
}

// a filter which will yeild all num which not devide by current prime filter
// that number will go for next prime filter ( bigger prime ) and will be pass
// as prime to be print in main when it pass all filter prime, then become new filter prime
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Printf("--> yeild %v by prime %v filter\n", i, prime)
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}
