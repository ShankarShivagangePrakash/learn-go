package main

import "fmt"

func main() {
	// Variable declaration with explicit type
	var age int = 25
	var name string = "Alice"

	fmt.Println("Age:", age)
	fmt.Println("Name1:", name)

	// shorthand declaration operator for creating variables
	a := 30
	b := 20
	c := a + b

	fmt.Println("Sum created using shorthand declaration operator:", c)

	// Constants
	const pi float32 = 3.14
	fmt.Println("Value of Pi:", pi)

	// Ouput functions
	fmt.Printf("Formatted Output: Name is %s and Age is %d\n", name, age) // print with formatting
	fmt.Println("Using Println: Name is", name, "and Age is", age)        // print with new line
	fmt.Print("Using Print: Name is ", name, " and Age is \n", age)         // print without new line

	var var1 int
	var var2 string

	fmt.Scan(&var1, &var2)
	fmt.Println("Scanned Values - var1:", var1, ", var2:", var2)

}
