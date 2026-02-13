package main

import (
	"fmt"
	"sync"
)

/*Function that will be run as a goroutine which accepts a WaitGroup pointer
* means it will receive the address of a WaitGroup variable */
func output(message string, wg *sync.WaitGroup) {
	fmt.Println(message)
	// indicate that the goroutine is done
	wg.Done()
}

func main() {

	// Declare a WaitGroup
	var wg sync.WaitGroup

	// Indicate that we are going to wait for 1 goroutine
	wg.Add(10)

	for i := 0; i < 10; i++ {

		/* start a new goroutine - thread and pass the WaitGroup pointer
		   & means we are passing the address of the WaitGroup variable */
		go output(fmt.Sprintf("Hello from a goroutine from go rountine number: %d" , i), &wg)

	}

	// Wait for all goroutines to finish
	wg.Wait()

	// main thread resumes execution after all goroutines are done - wg.wait() will be blocked until then
	fmt.Println("Hi from Main program")

	
}
