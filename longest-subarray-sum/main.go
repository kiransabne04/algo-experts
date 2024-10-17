package main

import "fmt"

func main() {
	val := LongestSubArrayWithSum([]int{1, 2, 3, 4, 3, 3, 1, 2, 1, 2}, 10)
	fmt.Println(val)
}

func LongestSubArrayWithSum(arr []int, k int) []int {
	start := 0
	maxLength := 0
	currentSum := 0
	res := [2]int{}

	for i := 0; i < len(arr); i++ {
		fmt.Println("->", start, i, arr[i])
		currentSum += arr[i]

		// checking for currentSum > k
		for currentSum > k && start <= i {
			//maxLength = max(maxLength, i-start+1)
			currentSum -= arr[start]
			start++

		}
		fmt.Println("currentsum, maxlen, i ", currentSum, k, maxLength, i, start)
		if currentSum == k && maxLength < i-start+1 {

			maxLength = i - start + 1
			res[0] = start
			res[1] = i
		}
	}
	return res[:]
}
