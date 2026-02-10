package main

import (
	"fmt"
)

func main() {

	// uses of slice
	// 1 => dynamic array , can increase the size of the array automatically
	// the coopy of the previous memory location will not be deleted permanently this is the only problem while using slices in go
	// 2 => it can also be used for sclicing the array

	// 1. syntax => var sclice_name = []datatype{} . here if we dinot mention the size then it is a sclice . If we mention the size then it it is a arruy

	var sclice = []int32{}
	sclice = append(sclice, 1)
	sclice = append(sclice, 2)
	sclice = append(sclice, 3)

	fmt.Println(sclice)

	// 2. syntax => using make(slice_name , no_of_intialValues , max_no_of_values) , there make is the keyword. we use infer(:=) keyword for this
	// here defineing the no_of_intialValues => means no fo empty values been created , so mention 0 => if u want to append the values from index0
	heros := make([]string,0)
	heros = append(heros, "ironman")
	heros = append(heros, "Spideman")
	heros = append(heros, "thor")
	heros = append(heros, "hulk") // automatically the size of the array is increased

	fmt.Println(heros)

	// 3 => can be used for slicing => syntax => array_name[start:end) -> start inclusive and the end is exclusive considered till end -1
	sliced_heros := heros[2:4] // here the 1 is inclusive and 4 is exclusive so it will consider the values from index 1 to index 3
	fmt.Println(sliced_heros)

	 arr1 := [6]int{10, 11, 12, 13, 14,15}
 	 myslice := arr1[2:4]
  	fmt.Println(myslice) // Output: [12 13]

}
