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
}

func sum(list []int) int {
	if len(list) < 1 {
		return 0
	}

	return list[0] + sum(list[1:])
}
