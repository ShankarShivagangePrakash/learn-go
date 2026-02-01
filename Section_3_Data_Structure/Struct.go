package main

import "fmt"

func main() {

	type Person struct {
		name string
		age int
		isMale bool
	}

	// structure object initialization method 1
	var p1 Person = Person{name: "Alice", age: 30, isMale: false}


	// structure object declaration
	var p2 Person

	// structure object initialization
	p2.name = "Bob"
	p2.age = 25
	p2.isMale = true

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
}


