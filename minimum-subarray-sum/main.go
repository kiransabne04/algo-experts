package main

import "fmt"

func main() {
	val := []int{2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := minSubarraySum(val)
	fmt.Println("res -> ", res)
}

func minSubarraySum(arr []int) int {
	globalMin := arr[0]
	localMin := arr[0]

	for i := 1; i < len(arr); i++ {
		localMin += min(localMin, localMin+arr[i])
		globalMin = min(globalMin, localMin)
	}
	return globalMin
}
