package main

import "fmt"

func main() {
	//array := []int{-1, -3, -10, -1100, -1100, -1101, -1102, -9001}
	array := []int{1, 3, 10, 1100, 1100, 1101, 1102, 9001}
	res := IsMonotonic1(array)
	fmt.Println(res)
}

func IsMonotonic1(array []int) bool {
	//
	if len(array) <= 2 {
		return true
	}

	direction := array[1] - array[0]
	for i := 2; i < len(array); i++ {
		//fmt.Println("arr -> ", direction, array[i], array[i]-array[i-1])
		if direction == 0 {
			direction = array[i] - array[i-1]
			continue
		}

		if checkDirection(direction, array[i], array[i-1]) {
			return false
		}
	}

	return true
}

func checkDirection(direction, current, previous int) bool {
	difference := current - previous
	if direction > 0 {
		return difference < 0
	}
	return difference > 0
}

func IsMonotonic(array []int) bool {
	// Write your code here.
	increasingCase, decreasingCase := true, true
	for i := 1; i < len(array); i++ {
		//fmt.Println("arr ", array[i], array[i+1])
		if array[i] < array[i-1] {
			fmt.Println(array[i], array[i+1])
			decreasingCase = false
		}
		if array[i] > array[i-1] {
			increasingCase = false
		}
	}

	return increasingCase || decreasingCase
}
