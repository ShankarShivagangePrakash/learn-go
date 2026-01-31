/*
 Declares the package name; package main is special—it defines an executable program (not a library).
 file containing main function shoule be in package main. Starting point of execution.
*/
package main           

import "fmt"           // Imports the "fmt" package, which contains formatting and I/O functions (like printing to the console).

func main() {          // Defines the main function—this is the entry point for the program; execution starts here.
    fmt.Println("Hello, World!") // Calls the Println function from the fmt package to print "Hello, World!" to the console.
}