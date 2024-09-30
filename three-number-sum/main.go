package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	targetSum := 0
	res := ThreeNumberSum(array, targetSum)
	fmt.Println("res -> ", res)
}

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	sort.Ints(array) // sort the array first
	triplets := [][]int{}
	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1
		for left < right {
			if array[i] + array[left] + array[right] == target {
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if array[i] + array[left] + array[right] < target {
				left += 1
			} else if array[i] + array[left] + array[right] > target {
				right -=1
			}
 		}
	}
	
	return triplets
}
