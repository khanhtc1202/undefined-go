package main

func main() {
    done := make(chan int)
    go func(){
		for i := 0; i < 10; i++ {
			println(i, "Hello World")
			done <- 1
		}
    }()
    <-done
    close(done)
}

