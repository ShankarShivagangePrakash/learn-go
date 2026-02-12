package main

import "fmt"

// Define an interface
type Greeter interface {
	Greet() string
}

// Declare child class which implements the interface
type EnglishGreeter struct{}

// Implement the interface method
// Synatx: func (receiver Type) MethodName() ReturnType { ... }
func (E EnglishGreeter) Greet() string {
	return "Hello!"
}

type SpanishGreeter struct{}

func (S SpanishGreeter) Greet() string {
	return "¡Hola!"
}

// dynamic binding - run-time polymorphism
func greetSomeone(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	eng := EnglishGreeter{}
	span := SpanishGreeter{}

	greetSomeone(eng)
	greetSomeone(span)
}
