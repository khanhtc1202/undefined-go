package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			println("Hello World")
			wg.Done()
		}
	}()

	// for loop inside forked go routine run faster than teardown prog in main routine ( after wg.Wait() statement )
	// => wg.Done() be revoked one more time after the println statement be called => panic
	wg.Wait()
}

