package main

import (
	"fmt"
)


func main() {
	// start a new goroutine - thread using an anonymous function
	go func(message string) {
		fmt.Println(message)
	} ("Hello from a goroutine!")

	// main thread continues execution concurrently - does not wait for the goroutine to finish
	fmt.Println("Hi from Main program")
}
