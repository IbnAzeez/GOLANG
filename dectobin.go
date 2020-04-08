package main

import ("fmt"
	"strconv"
)

func main() {
    //reading an integer
    var decimal int
	fmt.Println("Please input the Decimal number")
    fmt.Scanln(&decimal)
	var decimal_no int64 = int64(decimal) 
	
	fmt.Println("The equivalent binary conversion is: ")
	fmt.Println(strconv.FormatInt(decimal_no, 2)) // 1111011
}