package main

import "fmt"

func main() {

	// map using var keyword
	var map1 = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
    fmt.Println("Map using var keyword:", map1)

	// map using make function
	map2 := make(map[int]string)
	map2[4] = "Four"
	map2[5] = "Five"
	map2[6] = "Six"
	fmt.Println("Map using make function:", map2)

	// remove map entry
	delete(map1, 2)
}
