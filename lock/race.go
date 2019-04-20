package main

func raceCondition() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		println("value is", data)
	}
}

func main() {
	raceCondition()
}
