package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	list := rand.Perm(50)
	fmt.Println("Generated list:", list)
	fmt.Println("Sum:", sum(list))
	fmt.Println("Count:", count(list))
	fmt.Println(maxNumber(list))
}

func sum(list []int) int {
	if len(list) < 1 {
		return 0
	}

	return list[0] + sum(list[1:])
}

func count(list []int) int {
	if len(list) < 1 {
		return 0
	}
	return 1 + count(list[1:])
}

func maxNumber(list []int) int {
	switch {
	case len(list) == 0:
		return 0
	case len(list) == 1:
		return list[0]
	case len(list) == 2:
		switch {
		case list[0] > list[1]:
			return list[0]
		default:
			return list[1]
		}
	}
	subMax := maxNumber(list[1:])
	if list[0] > subMax {
		return list[0]
	}
	return subMax
}
