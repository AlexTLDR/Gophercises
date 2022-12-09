package main

/*
Create a program which displays random moods, using a multi-dimensional array. When the program is run with
name and positive/negative, it prints the name and the mood from the array.
Ex: run go main.go alex positive
alex feels cheerful
run go main.go alex negative
alex feels sad
*/

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	moods := [2][6]string{
		{"sad", "terrible", "gloomy", "reflective", "depressed", "unhappy"},
		{"happy", "good", "awesome", "jovial", "friendly", "cheerful"},
	}

	if len(os.Args) != 3 {
		fmt.Println("[name] [positive|negative]")
		return
	}

	name := os.Args[1]

	/*
		switch solution

	switch os.Args[2]{
	case "positive":
		n := rand.Intn(len(moods[1]))
		fmt.Printf("%s feels %s!\n", name, moods[1][n])
	case "negative":
		n := rand.Intn(len(moods[0]))
		fmt.Printf("%s feels %s!\n", name, moods[0][n])

	}

	 */

	n := rand.Intn(len(moods[0]))
	var i int // mood indicator
	if os.Args[2] != "negative" {
		i = 1
	}
	fmt.Printf("%s feels %s!\n", name, moods[i][n])


}


