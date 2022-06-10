package main

import (
	"fmt"
	"math"
)

func main() {
	res := SortedSquaredArray([]int{1, 2, 3, 5, 6, 8, 9})
	fmt.Println("res -> ", res)
}

func SortedSquaredArray(array []int) []int {
	ans := make([]int, len(array))
	sIndex := 0
	lIndex := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {
		smallVal := array[sIndex]
		largeVal := array[lIndex]

		if math.Abs(float64(smallVal)) > math.Abs(float64(largeVal)) {
			ans[i] = smallVal * smallVal
			sIndex++
		} else {
			ans[i] = largeVal * largeVal
			lIndex--
		}
	}

	return ans
}
