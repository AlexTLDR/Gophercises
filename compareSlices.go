package main

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//  4. Print the statement with a variable for not equal
//
// EXPECTED OUTPUT
//
//   They are equal.
//
// namesA := "Da Vinci, Wozniak, Carmack"
// namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
//
// ---------------------------------------------------------

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	sNamesA := strings.Split(namesA, ", ")

	sort.Strings(namesB)
	sort.Strings(sNamesA)

	var ok string
	if len(namesB) != len(sNamesA) {
		ok = "not "

	}
	for i, name := range namesB {
		if name != sNamesA[i] {
			ok = "not "
			break
		}
	}
	fmt.Printf("They are %sequal\n", ok)
}
