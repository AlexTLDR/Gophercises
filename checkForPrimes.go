// ---------------------------------------------------------
// EXERCISE: Check if prime number
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------
package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
skip:
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			continue skip
		} else if big.NewInt(int64(num)).ProbablyPrime(0) {
			fmt.Println(num, "is prime")
		}
	}

}
