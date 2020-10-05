package main

// ---------------------------------------------------------
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go aliens
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 8.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	s := "'s richter scale is"

	switch m := strings.Join(os.Args[1:], " ") ;m {
	case "micro":
fmt.Printf("%s%s less than 2.0\n",m, s)
	case "very minor":
		fmt.Printf("%s%s 2-2.9\n",m, s)
	case "minor":
		fmt.Printf("%s%s 3-3.9\n",m, s)
	case "light":
		fmt.Printf("%s%s 4-4.9\n",m, s)
	case "moderate":
		fmt.Printf("%s%s 5-5.9\n",m, s)
	case "strong":
		fmt.Printf("%s%s 6-6.9\n",m, s)
	case "major":
		fmt.Printf("%s%s 7-7.9\n",m, s)
	case "great":
		fmt.Printf("%s%s 8-8.9\n",m, s)
	case "massive":
		fmt.Printf("%s%s 10+\n",m, s)
	default:
		fmt.Printf("%s%s unknown\n",m, s)
	}
}
