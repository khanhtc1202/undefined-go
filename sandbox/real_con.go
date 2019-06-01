package main

import (
	"sync"
	"time"
)

type Container struct {
	fork   int
	member IMember
}

func (c *Container) SleepMulti() {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < c.fork; i++ {
		wg.Add(1)
		go func(time int, wg *sync.WaitGroup) {
			println("Go to sleep", time)
			println("Member info:", c.member)
			c.member.Sleep()
			println("Awake", time)
			defer wg.Done()
		}(i, &wg)
	}

	wg.Wait()
	finish := time.Now()
	println("Evg:", finish.Sub(start).String())
}

type IMember interface {
	Sleep()
}

type Member struct {
	sleepTime int
}

func (m *Member) Sleep() {
	time.Sleep(time.Duration(m.sleepTime) * time.Second)
}

func main() {
	member := Member{sleepTime: 1}
	container := Container{fork: 4, member: &member}

	println("Go to sleep multi")
	container.SleepMulti()
}
