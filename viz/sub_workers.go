package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	WORKERS    = 5
	SUBWORKERS = 3
	TASKS      = 20
	SUBTASKS   = 10
)

func subWorker(subtasks chan int) {
	for {
		task, ok := <-subtasks
		if !ok {
			return
		}
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Println(task)
	}
}

func hasSubWorker(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			return
		}

		subTasks := make(chan int)
		for i := 0; i < SUBWORKERS; i++ {
			go subWorker(subTasks)
		}
		for i := 0; i < SUBTASKS; i++ {
			task1 := task * i
			subTasks <- task1
		}
		close(subTasks)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(WORKERS)
	tasks := make(chan int)

	for i := 0; i < WORKERS; i++ {
		go hasSubWorker(tasks, &wg)
	}

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()
}
