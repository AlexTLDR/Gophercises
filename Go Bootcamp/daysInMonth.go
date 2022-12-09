package main

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "december" has 31 days.
//
//  go run main.go dEcEmBeR
//    "december" has 31 days.
// ---------------------------------------------------------


import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	currentTime := time.Now()
	currentYear := currentTime.Year()

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}
	arg := strings.ToLower(os.Args[1])

	if arg == "january" || arg == "march" || arg == "may" || arg == "july" || arg == "august" || arg == "october" || arg == "december" {
		fmt.Printf("%q has 31 days.\n", arg)
	} else if arg == "april" || arg == "june" || arg == "september" || arg == "november" {
		fmt.Printf("%q has 30 days.\n", arg)
	} else if ((currentYear%4 == 0 && currentYear%100 != 0) || currentYear%400 == 0 ) && arg == "february" {
		fmt.Printf("%q has 29 days.\n", arg)
	} else if arg == "february" {
		fmt.Printf("%q has 28 days.\n", arg)
	} else {
		fmt.Printf("%q is not a month", arg)
	}

}
