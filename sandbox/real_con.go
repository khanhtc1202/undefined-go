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

type HaveLockMember struct {
	lock      sync.Mutex
	sleepTime int
}

func (m *HaveLockMember) Sleep() {
	m.lock.Lock()
	time.Sleep(time.Duration(m.sleepTime) * time.Second)
	m.lock.Unlock()
}

func main() {
	member := Member{sleepTime: 1}
	container := Container{fork: 4, member: &member}

	println("Go to sleep multi without lock")
	container.SleepMulti()

	println("------------")
	haveLockMember := HaveLockMember{sleepTime: 1}
	container1 := Container{fork: 4, member: &haveLockMember}

	println("Go to sleep with lock")
	container1.SleepMulti()
}
