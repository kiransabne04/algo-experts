package main

import "fmt"

func main() {
	val := LongestSubarrayWithSum([]int{61, 54, 1, 499, 2212, 4059, 1, 2, 3, 1, 3}, 19) // 4, 8
	fmt.Println(val)
}

func LongestSubarrayWithSum(array []int, targetSum int) []int {
	// Write your code here.
	res := make([]int, 2)
	start := 0
	//end := len(array) - 1
	currentSum := 0
	maxLength := 0

	if len(array) == 0 {
		res[0] = 0
		res[1] = 0
		return res[:]
	}

	for end := 0; end < len(array); end++ {
		fmt.Println(array[start], array[end])
		currentSum += array[end]
		// if currentSum > targetsum then shrink the window from start pointer
		for currentSum > targetSum && start <= end {
			currentSum -= array[start]
			start++
		}

		// if currentSum matches the targetsum
		if currentSum == targetSum && end-start+1 > maxLength {
			maxLength = end - start + 1
			res[0] = start
			res[1] = end
		} else {
			return []int{}
		}
	}

	if maxLength == 0 {
		return []int{}
	}
	return res[:]

}
