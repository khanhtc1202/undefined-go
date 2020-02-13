package main

import (
	"log"
	"time"
)

func worker(queue chan int, workernumber int, done chan bool, killsignal chan bool) {
	for {
		select {
		case k := <-queue:
			log.Print("worker ", workernumber, " done with job ", k)
			done <- true
		case <-killsignal:
			log.Print("worker halted, number", workernumber)
			return
		}
	}
}

func main() {
	kill := make(chan bool)
	q := make(chan int)
	done := make(chan bool)

	numberOfWorkers := 4
	for i := 0; i < numberOfWorkers; i++ {
		go worker(q, i, done, kill) // waiting for jobs
	}

	numberOfJobs := 17
	for j := 0; j < numberOfJobs; j++ {
		go func(j int) {
			q <- j // pass job to queue of jobs
		}(j)
	}

	// wait until all of jobs done
	for c := 0; c < numberOfJobs; c++ {
		<-done
	}

	close(kill)
	// we need this to make sure program wait until kill signal be passed
	// NOTE that defer not work here since the program can be halted before defer block done
	// otherwise use waitgroup if you don't want to use sleep here
	time.Sleep(2 * time.Second) 
}
