package main

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println(`[command] [string]

Available commands: lower, upper and title`)
		return
	}
	arg := strings.Join(os.Args[2:]," ")
	switch cmd := os.Args[1];cmd{
	case "lower":
		fmt.Println(strings.ToLower(arg))
	case "upper":
		fmt.Println(strings.ToUpper(arg))
	case "title":
		fmt.Println(strings.Title(arg))
	default:
		fmt.Printf("Unknown command: %q\n", os.Args[1])
	}
}
