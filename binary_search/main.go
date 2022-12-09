package main

import "fmt"

func main() {
	list := []int{}
	for i := 0; i < 101; i++ {
		list = append(list, i)
	}
	fmt.Println(binarySearch(21, list))
}

func binarySearch(i int, list []int) bool {
	low := list[0]
	high := list[len(list)-1]

	for low <= high {
		mid := (low + high) / 2
		check := list[mid]
		if check == i {
			return true
		} else if check > i {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
