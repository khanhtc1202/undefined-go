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
    // below `close` does not have meaning here, since for loop still run inside forked 
    // routine but program will stop right after println statment be executed
    close(done)
}

