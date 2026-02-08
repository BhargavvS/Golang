package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// goal take the user rating and displace a custom message

	reader := bufio.NewReader(os.Stdin)

	//frontend

	// steps:
	// 1. take the user name
	// 2. take the rating

	var username string
	fmt.Println("Enter your full name")
	username, _ = reader.ReadString('\n')

	fmt.Println("Enter your user expirence rating between 1 to 5: ")
	myRating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)

	// backend
	// steps:
	// 1. print name , rating and the time stamp

	
	fmt.Printf("Thank you %v for your rating. \n You have successfully submitted the rating %v out of 5 at %v", username, mynumRating, time.Now().Format(time.Stamp))

}
