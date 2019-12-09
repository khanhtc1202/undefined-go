package main

import "fmt"

func main() {
    done := make(chan int, 1)

    go func(){
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World")
		done <- 1
	}
    }()

    <-done
}

