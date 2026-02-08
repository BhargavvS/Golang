package main

import (
	"fmt"
)

func main() {
	// two ways

	// 1=> first declare the array , syntax => var arrayName [size(mandatory)]dataType
	// the assign the values using indexes

	var names [3]string
	names[0] = "ironman"
	names[1] = "hulk"
	names[2] = "krishna"

	fmt.Println(names) // [ironman hulk krishna] => space separated

	// 2 => initialization of the array while declaring it
	// syntax => var arrayname = [size(mandatory)]datatype{val1 , val2 , val3}

	var superHeros = [3]string{"krishna" , "rama" ,"govinda"}
	fmt.Println(superHeros) // [krishna rama govinda]

}