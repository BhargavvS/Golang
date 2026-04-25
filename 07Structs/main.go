package main

import "fmt"

func main() {
	// an array without size is an slices

	// 1st way to declare an slice
	// var dynamicArray1 []int
	// dynamicArray1 = append(dynamicArray1, 1)
	// dynamicArray1 = append(dynamicArray1, 2)

	// fmt.Println(dynamicArray1)
	// fmt.Println(len(dynamicArray1))

	// @nd method is using make keyword
	// make(type of slice , initial size , capacity)

	// capacity -> the max si8ze of the slice initailly , if we add more elemennts to the slice beyound 	the    capacity , the capacity is automatically increses
	dynamicarray2 := make([]int, 0, 10)
	dynamicarray2 = append(dynamicarray2, 1)
	dynamicarray2 = append(dynamicarray2, 2)

	fmt.Println("len = ", len(dynamicarray2))
	fmt.Println("cap =", cap(dynamicarray2))

	dynamicarray2 = append(dynamicarray2, 1)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)
	dynamicarray2 = append(dynamicarray2, 2)

	fmt.Println("len = ", len(dynamicarray2))
	fmt.Println("cap =", cap(dynamicarray2)) // now the cap change to 20 fron 10 =>  " newcap = 2 x cap "
}
