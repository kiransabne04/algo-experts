package main

import "fmt"

func main() {
	val := LongestSubarrayWithSum([]int{1, 2, 3, 4, 3, 3, 1, 2, 1}, 10) // 4, 8
	fmt.Println(val)
}

func LongestSubarrayWithSum(array []int, targetSum int) []int {
	// Write your code here.
	res := [2]int{}
	start := 0
	//end := len(array) - 1
	currentSum := 0
	maxLength := 0

	for end := 0; end < len(array); end++ {
		fmt.Println(array[start], array[end])
		currentSum += array[end]
		// if currentSum > targetsum then shrink the window from start pointer
		for currentSum > targetSum && start <= end {
			currentSum -= array[start]
			start++
		}

		// if currentSum matches the targetsum
		if currentSum == targetSum && end-start > maxLength {
			maxLength = end - start
			res[0] = start
			res[1] = end
		}
	}

	return res[:]

}
