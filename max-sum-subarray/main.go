package main

import "fmt"

func main(){
	array := []int{1, 2, 3, -2, 5}
	val := MaxProduct(array)
	fmt.Println(val)
}

func MaxSum(arr []int) int {
	globalMax := arr[0]
	localMax := arr[0]

	for i := 1; i < len(arr); i++ {
		localMax = max(arr[i], localMax + arr[i])

		globalMax = max(globalMax, localMax)
		fmt.Println(localMax, globalMax)

	}
	return globalMax
}

func MaxProduct(arr []int) int {
	globalMax := arr[0]
	localMax := arr[0]

	for i := 1; i < len(arr); i ++{
		localMax = max(arr[i], localMax * arr[i])

		globalMax = max(globalMax, localMax)
	}
	return globalMax
}

func MinSum(array []int) int {
	globalMin := array[0]
	localMin := array[0]

	for i := 1; i < len(array); i++ {
		localMin = min(array[i], localMin + array[i])
		globalMin = min(globalMin, localMin)
	}
	return globalMin
}

func MaxSumCircularSubarray(array []int) int {
	array = []int{1, -2, 3, -2}

	maxKadane := MaxSum(array)

	minKadane := MinSum(array)

	total := totalSum(array)

	if total == minKadane {
		// if all elements are negative then return 
		return maxKadane
	}

	return max(maxKadane, total-minKadane)
}

func totalSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}