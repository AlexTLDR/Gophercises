package main

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   From the string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
// ---------------------------------------------------------
//func main() {
// uncomment the declaration below
// data := "2 4 6 1 3 5"
//}
// ---------------------------------------------------------

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"
	//   1. Convert the string to an []int
	dataSlice := strings.Split(data, " ")
	/*
	ints := make([]int, len(dataSlice))
	for i, s := range dataSlice {
		ints[i], _ = strconv.Atoi(s)
	}
	 */
	//Refactor:
	var ints []int
	for _, s := range dataSlice {
		n, _ := strconv.Atoi(s)
		ints = append(ints, n)
	}
	//   2. Print the slice
	fmt.Printf("nums: %v\n", ints)
	//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
	evens := ints[:3]
	fmt.Printf("evens: %v\n", evens)
	//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
	odds := ints[3:]
	fmt.Printf("odds: %v\n", odds)
	//   5. Slice it for the two numbers at the middle
	middle := ints[2:4]
	fmt.Printf("middle: %v\n", middle)
	//   6. Slice it for the first two numbers
	first2 := ints[:2]
	fmt.Printf("first2: %v\n", first2)
	//   7. Slice it for the last two numbers (use the len function)
	last2 := ints[len(ints)-2:]
	fmt.Printf("last2: %v\n", last2)
	//   8. Slice the evens slice for the last one number
	evensLast := evens[len(evens)-1:]
	fmt.Printf("evens last: %v\n", evensLast)
	//   9. Slice the odds slice for the last two numbers
	oddsLast := odds[len(odds)-2:]
	fmt.Printf("odds last2: %v\n", oddsLast)
}
