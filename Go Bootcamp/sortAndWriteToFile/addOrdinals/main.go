package main

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file with their ordinals
//
//  Use the previous exercise.
//
//  This time, print the sorted items with ordinals
//  (see the expected output)
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     1. apple
//     2. banana
//     3. orange
// ---------------------------------------------------------

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)

	var sorted []byte
	for i, arg := range args {
		sorted = strconv.AppendInt(sorted, int64(i+1), 10)
		sorted = append(sorted, '.', ' ')
		sorted = append(sorted, arg...)
		sorted = append(sorted, '\n')
	}
	err := ioutil.WriteFile("sorted.txt", sorted, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
