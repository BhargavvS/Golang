package main

import "fmt"

func main() {

	age := 16

	// if age >= 18 {
	// 	fmt.Println("The person is an adult")
	// } else {
	// 	fmt.Println("The person is not an adult")
	// }

	if age >= 18 {
		fmt.Println("The person is an adult")
	} else if age >= 16 {
		fmt.Println("The person is an teenager")
	} else {
		fmt.Println("The person is a child")
	}

	// there is no turnary operator in golang
}
