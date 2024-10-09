package main

import (
	"fmt"
	"sort"
)

func main() {
	arrayOne := []int{-1, 5, 10, 20, 28, 3}
	arrayTwo := []int{26, 134, 135, 15, 17}
	val := SmallestDifference(arrayOne, arrayTwo)
	fmt.Println(val)
}

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	res := make([]int, 2)
	sort.Ints(array1)
	sort.Ints(array2)
	fmt.Println(array1, array2)
	//to do:: pending
	for i := 0; i < len(array1); i++ {
		firstPtr := 0
		secondPtr := 0
		if array1[firstPtr]-array2[secondPtr] == 0 {
			res[0] = array1[firstPtr]
			res[1] = array2[secondPtr]
		} else if array1[firstPtr] > array2[secondPtr] {
			secondPtr += 1
		}
	}

	return nil
}
