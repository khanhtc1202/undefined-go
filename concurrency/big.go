package main

import "sync"

const BIG = 20

func Worker(id int, ping <-chan int, done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case m := <- ping:
			println("receive", m, "from", id)
		case <- done:
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(BIG)

	ping := make(chan int)
	done := make(chan bool, BIG)
	for j := 0; j < BIG; j++ {
		go Worker(j, ping, done, &wg)
	}

	for i := 0; i < BIG; i++ {
		ping <- 1
	}
	for i := 0; i < BIG; i++ {
		ping <- 2
	}
	for i := 0; i < BIG; i++ {
		done <- true
	}

	wg.Wait()
	close(ping)
}
