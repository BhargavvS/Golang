package main

import "fmt"

func main() {
	start :=1

	for i:=0;i<10;i++{
		fmt.Println(i + start)
	}

	// loop through an array
	arr := []string{"apple", "banana", "watermelon"}

	for i:=00;i<len(arr);i++ {
		fmt.Println(arr[i])
	}

	//while loops in go => we donot have while loops in go but we can use for loop to achieve the same result

	for start < 100 {
		start += start

		if start == 64{
			break
		}
		
		fmt.Println(start)
	}

}