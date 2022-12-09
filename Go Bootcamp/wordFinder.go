package main

// ---------------------------------------------------------
// EXERCISE: Word Finder
//
//   Your goal is to search for the words inside the corpus.
//
//   1. Get the search query from the command-line (it can be multiple words)
//
//   2. Filter these words, do not search for them:
//      and, or, was, the, since, very
//
//      To do this, use an array for the filtered words.
//
//   3. Print the words found.
//
// RESTRICTION
//   + The search and the filtering should be case insensitive
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me a word to search.
//
//   go run main.go and was
//
//   go run main.go AND WAS
//
//   go run main.go cat beginning
//     #2 : "cat"
//     #11: "beginning"
//
//   go run main.go Cat Beginning
//     #2 : "cat"
//     #11: "beginning"
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	args := os.Args[1:]
	filteredWords := [...]string{"and", "or", "was", "the", "since", "very"}
	words := strings.Fields(corpus)

	if len(args) == 0 {
		fmt.Println("Please give me a word to search.")
	}

arguments:
	for _, q := range args {

		for i, w := range words {
			for _, f := range filteredWords {
				if strings.ToLower(q) == strings.ToLower(f) {
					continue arguments
				}
			}

			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue arguments
			}
		}
	}
}
