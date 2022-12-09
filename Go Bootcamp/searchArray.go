package main

// ---------------------------------------------------------
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	if len(os.Args) < 2 {
		fmt.Println("Tell me a book title")
		return
	}

	search := strings.Join(os.Args[1:], " ")

	fmt.Println("Search Results:")

var found bool
	for i := 0; i < len(books); i++ {
		if strings.Contains(strings.ToLower(books[i]), strings.ToLower(search)) {
			fmt.Printf("+ %s\n", books[i])
			found = true
		}
	}
	if !found {
		fmt.Printf("We don't have the book: %q\n", search)
	}
}
