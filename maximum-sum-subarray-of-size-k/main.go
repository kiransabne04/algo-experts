package main

import "fmt"

func main() {

	val := maximumSumSubarraySizeK([]int{2, 1, 5, 1, 3, 2}, 3)
	fmt.Println("val -> ", val)

}

func maximumSumSubarraySizeK(arr []int, k int) int {

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	fmt.Println(windowSum)
	maxSum := windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i]
		windowSum -= arr[i-k]
		fmt.Println(i, k, "i & k -> ", i, i-k)
		maxSum = max(maxSum, windowSum)
	}
	return maxSum
}
