package main

import "fmt"

// the code which is declared defer are moved to the bottom of the program and execute in reverse order of the declaration ie..(LIFO) is follwed while execution
func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}
