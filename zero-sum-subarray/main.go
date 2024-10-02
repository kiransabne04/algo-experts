package main

import "fmt"

func main(){
	array := []int{-5, -5, 2, 3, -2}
	val := ZeroSumSubarray(array)
	fmt.Println(val)
}


func ZeroSumSubarray(nums []int) bool {
	// Write your code here.
	prefixSum := 0
	prefixSumMap := make(map[int]int)
	for i := 0; i < len(nums); i++{
		prefixSum += nums[i]

		if prefixSum == 0 {
			return true
		}

		if _, exist := prefixSumMap[prefixSum]; exist {
			return true
		}

		prefixSumMap[prefixSum] = i
	}
	return false
}
