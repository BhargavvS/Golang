package main

import (
	"fmt"
)

func main() {
	// two ways to create pointer in go
	
	// 1=> syntax => ( var pointer_name *dataType )

	// example
	// var p *int

	// 2=> create a pointer to the variable created
	// example

	var score float64 = 99.2
	p := &score // here p stores the address of the pointer p and not the actual value

	fmt.Println(score)
	fmt.Println(p) // prints the address of the score
	
	score = score * 2.2 // both the pointer is updated automatically

	// print the actual value of the pointer then , use => *pointer_name
	fmt.Println(score)
	fmt.Println(*p)


	

}
