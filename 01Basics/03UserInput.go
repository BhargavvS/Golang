package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


func main() {
	// var myString string;
	// fmt.Scanln(&myString)
	// fmt.Println(myString)

	// var name string ="Bhargav"
	// var a , _ = fmt.Println(name) // here _ means i know there is some value but doesnot know what it is actually 
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")
	// mystring, _ := reader.ReadString('\n')
	// fmt.Println(mystring)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the user rating")
	myRating , _ := reader.ReadString('\n')
	mynumRating , _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64) // 64 => the bytes we are considering 
	fmt.Println(mynumRating + 2)

	// here we are using string.trimSpace => parsefloat also consider the \n and the space and 
	// it cannot be converted to number ,so we have to remove them first and then pass it 
}