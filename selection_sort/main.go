package main

import "fmt"

func main() {
	list := []int{78, 34, 90, 12, 53, 23, 65, 89, 89, 92, 90}
	fmt.Println(findSmallest(list))
	fmt.Println(list)
}

func findSmallest(list []int) int {
	smallest := list[0]
	for i := 1; i < len(list); i++ {
		if smallest > list[i] {
			smallest = list[i]
		}
	}
	return smallest
}
