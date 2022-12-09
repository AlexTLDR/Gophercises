package main

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Please provide min max integer values.")
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err != nil || err2 != nil || min >= max {
		fmt.Println("Wrong values")
		return
	}

	var (
		sum int
		i   = min
	)
	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}
		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}
		sum += i
		i++
	}
	fmt.Printf(" = %d\n", sum)

}
