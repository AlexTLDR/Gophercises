package main

import "fmt"

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

func main() {
	printString("hello")
	printInt(5)
	printBool(false)

	printGerneric("hello")
	printGerneric(5)
	printGerneric(false)
}

// func printGerneric[T string | int | bool](t T) { fmt.Println(t) }
func printGerneric[T any](t T) { fmt.Println(t) }
