package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {

	if len(os.Args) == 1 {
		fmt.Println("Input a letter!")
		os.Exit(2)
	}
	arg := strings.ToLower(os.Args[1])
	fmt.Println(arg)

	_, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err == nil {
		fmt.Println("First input parameter must be a letter")
		os.Exit(69)
	}

	if len(arg) != 1 {
		fmt.Println("Give me a letter!")
	} else if arg == "a" || arg == "e" || arg == "i" || arg == "o" || arg == "u" {
		fmt.Println(strconv.Quote(arg), "is a vowel.")
	} else if arg == "y" || arg == "w" {
		fmt.Println(strconv.Quote(arg), "is sometimes a vowel, sometimes not.")
	} else {
		fmt.Println(strconv.Quote(arg), "is a consonant.")
	}

}



