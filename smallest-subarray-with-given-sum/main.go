package main

import (
	"fmt"
	"math"
)

// Problem: Given an array of integers and a number S, find the length of the smallest subarray whose sum is greater than or equal to S.
// Approach: Sliding window with variable window size.
// Example:
// plaintext
// Copy code
// Input: arr = [2, 1, 5, 2, 3, 2], S = 7
// Output: 2
// Explanation: The smallest subarray with a sum of at least 7 is [5, 2].
func main() {
	val := smallestSubarrayGivenSum([]int{2, 1, 5, 2, 3, 2}, 7)
	fmt.Println(val)
}

func smallestSubarrayGivenSum(arr []int, k int) int {
	left := 0
	currentSum := 0
	minLength := math.MaxInt64

	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]

		fmt.Println("current Sum -> ", currentSum)

		for currentSum >= k {
			minLength = min(minLength, i-left+1)
			currentSum -= arr[left]
			left++
		}
		fmt.Println("current sum a-> ", currentSum, minLength, left)

	}

	return 0
}
