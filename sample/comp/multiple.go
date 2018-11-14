package comp

type MultiplePair struct {
	x, y int64
}

func SingleCore(m []MultiplePair) (sum int64) {
	for _, p := range m {
		r := p.x * p.y
		sum += r
	}
	return
}

func NonBufferChanGoRoutine(m []MultiplePair) (sum int64) {
	mul := make(chan int64)
	for _, p := range m {
		go func(p MultiplePair) {
			r := p.x * p.y
			mul <- r
		}(p)
	}

	counter := 0
	for {
		select {
		case r := <-mul:
			sum += r
			counter++
			if counter == len(m) {
				return
			}
		}
	}
}
