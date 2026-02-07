package main
import "fmt"

func main() {
// there are many way for decalrinng a variable in go

// basic syntax => var/const variableName dataType = value
var superman string = "I am a superman"
fmt.Println(superman)

var ironman string = "I am a ironman"
fmt.Println(ironman) 
// fmt.Printf("%v %T", ironman , ironman) // there %v => prints the value and %T => returns the datatype of the variable

// 2 => using "":=" this 
// using this operator it can automatically detect the type of the variable and user doesnot neeed to explicity mention it

spiderman := "I am an amazing spiderman"
fmt.Println(spiderman)
// fmt.Printf("%v %T" , spiderman , spiderman) // retuns => "I am an amazing spiderman"(value) ,  string(datatype)

// 3 => declaring a bunch of variables with one var keyword
// here all the varibales are standlone variables and no variable is connected to each other there 
var (
	thor = "I am a thunder god"
	kingdom = "ashguard"
	hulk = "I am an angry hulk"
	color = "Green"
)
fmt.Println(thor , kingdom , hulk , color)
}