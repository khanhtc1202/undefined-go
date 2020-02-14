package main

import (
	"fmt"
	"context"
)

func GenerateNaturalWithCtx(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2;; i++ {
			select {
			case <- ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

// a filter which will yeild all num which not devide by current prime filter
// that number will go for next prime filter ( bigger prime ) and will be pass
// as prime to be print in main when it pass all filter prime, then become new filter prime
func PrimeFilterWithCtx(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <- ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateNaturalWithCtx(ctx)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilterWithCtx(ctx, ch, prime)
	}
	cancel()
}
