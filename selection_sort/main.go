package main

import "fmt"

func main() {
	integers := []int{78, 34, 90, 12, 53, 23, 65, 89, 89, 92, 90}
	strings := []string{"zz", "tt", "yy", "bb", "dd", "aa", "ss", "qq"}
	selectionSortInts(integers)
	fmt.Println(integers)
	selectionSortStrings(strings)
	fmt.Println(strings)

}

func selectionSortInts(list []int) []int {
	for i := 0; i < len(list); i++ {
		minIdx := i
		for j := i; j < len(list); j++ {
			if list[j] < list[minIdx] {
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
	return list
}

func selectionSortStrings(list []string) []string {
	for i := 0; i < len(list); i++ {
		minIdx := i
		for j := i; j < len(list); j++ {
			if list[j] < list[minIdx] {
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
	return list
}
