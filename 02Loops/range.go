package main

import (
	"fmt";
)

func main() {

	// to loop through the arrays

	// arr := []int{1,2,3,4,5}


	// range over array => index , number can be accesed
	// for i, num := range arr {
	// 	fmt.Println("index:", i , "number: ",num)
	// }

	// range over maps

	// mpp := map[string]int{"APPLE": 150, "BANANA": 50, "ORANGE": 100}

	// for key, val := range mpp {
	// 	fmt.Println("key: ", key, "value: ", val)
	// }

	// Range over strings

	for index , val := range "go"{
		fmt.Println("index: ", index, "value: ", val , "string val: ", string(val))
	}

}