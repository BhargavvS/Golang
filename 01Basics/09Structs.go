package main

import (
	"fmt"
)

type Student struct {
	id   int
	Name string
	age  int
	sec  string

	// func display() {
	// 	fmt.Println("ID:", id)
	// 	fmt.Println("Name:", Name)
	// 	fmt.Println("Age:", age)
	// 	fmt.Println("Section:", sec)
	// }
}

func main() {
	var s1 = Student{1, "Bhargav", 23, "A"}
	fmt.Println(s1)
}
