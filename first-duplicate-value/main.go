package main

import (
	"fmt"
)

func main() {
	array := []int{2, 1, 5, 2, 3, 3, 4}
	val := FirstDuplicateValue(array)
	fmt.Println(val)
}

func FirstDuplicateValue(array []int) int {
	// Write your code here.
	seenMap := make(map[int]bool)
	for _, v := range array {
		if _, found := seenMap[v]; found {
			return v
		}
		seenMap[v] = true
	}
	return -1
}
