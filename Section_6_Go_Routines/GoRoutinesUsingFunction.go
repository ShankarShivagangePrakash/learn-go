package main

import (
	"fmt"
	"time"
)

func output(message string) {
	fmt.Println(message)
	// simulate a long-running task
	time.Sleep(10 * time.Second)
}

func main() {
	// start a new goroutine - thread
	go output("Hello from a goroutine!")

	// main thread continues execution concurrently - does not wait for the goroutine to finish
	fmt.Println("Hi from Main program")
}
