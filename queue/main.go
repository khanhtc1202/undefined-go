package main

import (
	"log"
)

func worker(queue chan int, workernumber int, done chan bool) {
	for j := range queue {
		log.Print("worker ", workernumber, " done with job ", j)
		done <- true
	}
}

func main() {
	q := make(chan int)
	done := make(chan bool)

	numberOfWorkers := 4
	for i := 0; i < numberOfWorkers; i++ {
		go worker(q, i, done) // waiting for jobs
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
}
