package main

import "fmt"

func main() {
	print()
	result := add(3,4)
	fmt.Println(result)
}

func print() {
	fmt.Println("Hello world")
}


// syntax => func func_name(variable1 datatype , variable2 datatype) return_type
func add(a int , b int) int {
	return a + b
}