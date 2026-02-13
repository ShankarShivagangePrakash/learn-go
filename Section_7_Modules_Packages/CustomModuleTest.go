package main

import (
	"fmt"
	// importing a custom package from local module
	"github.com/ShankarShivagangePrakash/learn-go/Section_7_Modules_Packages/Custom_Module"
)

func main() {
	area := Rectangle.Area(10.5, 5.5)
	fmt.Printf("Area of Rectangle: %.2f\n", area)
}
