package main

import "fmt"

func main() {
	array := []int{-5, -5, 2, 3, -2}
	val := ZeroSumSubarray(array)
	fmt.Println(val)
}

func ZeroSumSubarray(arr []int) bool {
	runningSum := 0
	seenSumIndex := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		runningSum += arr[i]
		fmt.Println("running sum", runningSum)
		if runningSum == 0 {
			return true
		}

		// check if running sum is found in map
		if _, found := seenSumIndex[runningSum]; found {
			fmt.Println("found the required value ", runningSum, "we are looking for in the map")
			return true
		}

		seenSumIndex[runningSum] = i
	}
	fmt.Println(seenSumIndex)

	return false

}
