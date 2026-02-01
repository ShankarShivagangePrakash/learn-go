package main

import "fmt"

func main() {
	// array using shorthand declaration operator
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array:", arr)

	// array using var keyword
	var arr1 [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array 1:", arr1)

	// array declaration
	var arr2 [3]int
	
	arr2[0] = 10
	arr2[1] = 20
	arr2[2] = 30
	fmt.Println("Array 2:", arr2)
}