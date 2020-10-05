package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if num := os.Args; len(os.Args) < 2 {
		fmt.Println("Pick a number")
		return
	} else if len(os.Args) >2 {
		fmt.Println("Too many arguments")
	} else if n, err :=strconv.Atoi(num[1]); err != nil{
		fmt.Printf("%q is not a number\n", num[1])
	} else if n%8 == 0 {
		fmt.Println(n, "is an even number and it's divisible by 8")
	} else if n%2 == 0 {
		fmt.Println(n, "is an even number")
	} else{
		fmt.Println(n, "is an odd number")
	}
}
