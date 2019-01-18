package main

import "fmt"

func main() {
	func() {
		fmt.Println("Box 1")
		fmt.Println("Box 2")
		fmt.Println("Box 3")
		fmt.Println("Box 4")
		fmt.Println("Box 5")
		fmt.Println("Box 6")
		fmt.Println("Box 7")
		fmt.Println("Box 8")
	}()
	fmt.Println("Box 9")
	fmt.Println("Box 10")
}
