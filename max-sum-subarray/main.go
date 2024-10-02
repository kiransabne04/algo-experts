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

func MaxSumCircularSubarray(array []int) int {
	array = []int{1, -2, 3, -2}

	globalMax := array[0]
	localMax := array[0]
	
}