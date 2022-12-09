package main

// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 8.92 EUR
//     10.50 USD is 8.09 GBP
//     10.50 USD is 1108.07 JPY
//
//   go run main.go 1
//     1.00 USD is 0.85 EUR
//     1.00 USD is 0.77 GBP
//     1.00 USD is 105.53 JPY
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const (
		EUR = iota
		GPB
		JPY
	)

	rates := [...]float64{
		EUR: 0.85,
		GPB: 0.77,
		JPY: 105.53,
	}

	if len(os.Args) != 2 {
		fmt.Println("Please provide the amount to be converted")
		return
	}
	amount := os.Args[1]
	if amount, err := strconv.ParseFloat(amount,64); err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	} else {
		fmt.Printf("%.2f USD is %.2f EUR\n", amount,rates[EUR] * amount)
		fmt.Printf("%.2f USD is %.2f GPB\n", amount,rates[GPB] * amount)
		fmt.Printf("%.2f USD is %.2f JPY\n", amount,rates[JPY] * amount)
	}
}
