package main

import (
	"fmt"
)

func main() {
	// we can use make and new keyword for the intialization
	// 1. new => it allocates the meemory but doesnot init the memoy
	// 2. make => it allocates the memory and also initials the values of the memeory , it inits the values with "0"

	// first using new keyword
	// syntax =>var mymap map[key_datatype]value_datatype 

	// var mymap map[string]int
	// mymap["krithik"] = 10
	// mymap["bhargav" ] = 100 

	// fmt.Println(mymap) // this will give the error because we have not intialized the map using new or make keyword

	names := make(map[string]int)

	names["krithi"] = 10
	names["bhargav"] = 100

	fmt.Println(names)

	score := names["krithi"]
	fmt.Println(score)

	// delete(names , "krithi")
	fmt.Println(names)

	for k,v := range names {
		fmt.Printf("The score of %v is %v \n", k ,v)
	}
}