package main

import (
	"fmt"
	"math"
)

func main() {
	val := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})
	fmt.Println(val)
}

func SubarraySort(array []int) []int {
	// Write your code here.
	res := []int{-1, -1}

	leftIdx := -1
	for i := 0; i < len(array) -1 ; i++ {
		if array[i] > array[i + 1] {
			leftIdx = i
			break
		}
	}

	rightIdx := -1
	for i := len(array) - 1; i > 0; i-- {
		fmt.Println("rightIdx ", array[i] , array[i-1])
		if array[i] < array[i-1]{
			rightIdx = i
			break
		}
	}

	fmt.Println("leftIdx & rightIdx -> ", leftIdx, rightIdx)

	if leftIdx == -1 {
		return res
	}

	minOutOrder := math.MaxInt64
	maxOutOrder := math.MinInt64

	for i := leftIdx; i <= rightIdx; i ++ {
		minOutOrder = min(minOutOrder, array[i])
		maxOutOrder = max(maxOutOrder, array[i])
	}

	fmt.Println("minOutOrder & maxOutOrder -> ", minOutOrder, maxOutOrder)

	// expanding to left side and check if any out of order is left
	for leftIdx > 0 && array[leftIdx-1] > minOutOrder {
		leftIdx--
	}

	// same for right index
	for rightIdx < len(array) - 1 && array[rightIdx + 1] < maxOutOrder {
		rightIdx++
	}

	res[0] = leftIdx
	res[1] = rightIdx

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}