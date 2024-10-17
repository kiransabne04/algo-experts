package main

import "fmt"

//Given an array of integers and a number K, find the maximum sum of a subarray of size

// Input: arr = [2, 1, 5, 1, 3, 2], K = 3
// Output: 9
// Explanation: Subarray with maximum sum is [5, 1, 3]

func main() {
    val := MaxSumSubArraySizeK([]int{2, 1, 5, 1, 3, 2}, 3)
    fmt.Println("val -> ", val)
}

func MaxSumSubArraySizeK(arr []int, k int) int {
	windowSum := 0
    maxSum := 0
    for i := 0; i < k; i++{
        windowSum += arr[i]
    }
    fmt.Println(windowSum)
    maxSum = windowSum

    for i := k; i < len(arr); i++ {
        fmt.Println("i & i-k", i, i-k)
        windowSum += arr[i] - arr[i-k]

        maxSum = max(maxSum, windowSum)
    }
    return maxSum
}