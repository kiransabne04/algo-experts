package main

import "fmt"

func main() {
	array := []int{5, 1, 4, 2}
	val := ArrayOfProducts1(array)
	fmt.Println(val)
}

func ArrayOfProducts(array []int) []int {
	// Write your code here.
	res := make([]int, len(array))
	//8, 40, 10, 20
	for i := range array {
		runningLength := 1
		for j := 0; j < len(array); j++ {
			if i != j {
				runningLength *= array[j]
				//fmt.Println("i -> ", i, j, runningLength)
			}
		}
		res[i] = runningLength
	}
	return res
}

func ArrayOfProducts1(array []int) []int {
	res := make([]int, len(array))
	right := make([]int, len(array))
	for i := range res {
		res[i] = 1
		right[i] = 1
	}
	leftRunningProduct := 1
	for i := range array {
		res[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}
	//fmt.Println("left -> ", res)

	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		res[i] *= rightRunningProduct
		rightRunningProduct *= array[i]
	}
	//fmt.Println("right -> ", res)
	return res
}

// 1
