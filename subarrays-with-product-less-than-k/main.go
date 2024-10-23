package main

import "fmt"

// Subarrays with Product Less Than K
// Problem: Given an array of positive integers and a number k, find the number of contiguous subarrays where the product of all the elements is strictly less than k.
// Approach: Sliding window with dynamic product calculation.
// Example:
// plaintext
// Copy code
// Input: nums = [10, 5, 2, 6], k = 100
// Output: 8
// Explanation: The subarrays with products less than 100 are:
// [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].

func main() {
	val := subarraysWithProducts([]int{10, 5, 2, 6}, 100)
	fmt.Println(val)
}

func subarraysWithProducts(arr []int, k int) int {
	start := 0
	product := 1
	count := 0

	for i := 0; i < len(arr); i++ {
		product = product * arr[i]

		for product >= k && start <= i {
			product = product / arr[start]
			start++
		}
		count += i - start + 1
	}

	return count
}
