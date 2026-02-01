package main

import "fmt"

func main() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	// create slice from array
	var slice []int = arr[0:len(arr)] // all elements of array

	fmt.Println("Array:", arr)
	fmt.Println("Slice using arrays:", slice)

	// slice using [] operator
	var slice2 []int = []int{10, 20, 30, 40, 50}
	fmt.Println("Slice using [] operator:", slice2)

	// slice using make function
	slice3 := make([]int, 3)

	slice3[0] = 100
	slice3[1] = 200
	slice3[2] = 300

	fmt.Println("Slice using make function:", slice3)


}