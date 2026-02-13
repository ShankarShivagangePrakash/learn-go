package main 

import "fmt"

	func main() {
	ch1 := make (chan int)
	ch2 := make (chan string)

	go func() {
		ch1 <- 42
		close(ch1)
		// ch1 <- 100 // panic: send on closed channel
	}()

	go func() {
		ch2 <- "Hello, Channels!"
	}()

	val1 := <- ch1
	val2 := <- ch2

	fmt.Println("Received from channel 1:", val1)
	fmt.Println("Received from channel 2:", val2)

}