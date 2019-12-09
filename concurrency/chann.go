package main

func main() {
    done := make(chan int, 1)

    go func(){
	for i := 0; i < 10; i++ {
		println("Hello World")
		done <- 1
	}
    }()

    <-done
}

