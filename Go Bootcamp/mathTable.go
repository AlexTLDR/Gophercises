package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Please provide the size of the table by using an operand (+,-,x,/,%) and an integer!")
		return
	}



	if max, err := strconv.Atoi(os.Args[2]); err != nil {
		fmt.Printf("%q is not an integer", max)
		return
	} else if os.Args[1] == "x" {

		fmt.Printf("%5s", "X")
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", j*i)
			}
			fmt.Println()
		}

	}else if os.Args[1] == "+" {

		fmt.Printf("%5s", "X")
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", j+i)
			}
			fmt.Println()
		}

	}else if os.Args[1] == "-" {

		fmt.Printf("%5s", "X")
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", j-i)
			}
			fmt.Println()
		}

	}else if os.Args[1] == "/"{
		fmt.Println("Because division doesn't work with 0, the 0 branch has been removed:")
		fmt.Printf("%5s", "X")
		for i := 1; i <= max; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 1; i <= max; i++ {
			fmt.Printf("%5d", i)
			for j := 1; j <= max; j++ {
				fmt.Printf("%5d", j/i)
			}
			fmt.Println()
		}

	}else if os.Args[1] == "%"{
		fmt.Printf("%5s", "X")
		for i := 1; i <= max; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 1; i <= max; i++ {
			fmt.Printf("%5d", i)
			for j := 1; j <= max; j++ {
				fmt.Printf("%5d", j%i)
			}
			fmt.Println()
		}

	} else {
		fmt.Printf("%q is not an operand\n", os.Args[1])
	}
}
