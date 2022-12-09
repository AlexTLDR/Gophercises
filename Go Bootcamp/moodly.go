package main

/*
Create a program which displays random moods, using an array.
Ex: run go main.go alex
alex feels cheerful
 */

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	moods := [...]string{
		"sad",
		"terrible",
		"happy",
		"good",
		"awesome",
		"jovial",
		"friendly",
		"cheerful",
		"reflective",
		"gloomy",
	}

	if len(os.Args) != 2 {
		fmt.Println("Provide name as argument")
		return
	}
	name := os.Args[1]
	n := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[n])
}
