package main

import "fmt"

func main() {
	array := []int{-1, -100, -10, -1100, -1100, -1101, -1102, -9001}
	val := IsMonotonic(array)
	fmt.Println(val)
}


func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	// Write your code here.
	direction := array[1] - array[0]
	fmt.Println("first_direction::", direction)
	for i := 2; i < len(array); i++ {
		if direction == 0 {
			direction = array[i] - array[i-1]
			continue
		}
		fmt.Println("loop_array::", array[i], array[i-1], array[i] - array[i-1])

		if checkDirection(direction, array[i], array[i-1]) {
			return false
		}
	}
	return true
}

func checkDirection(direction, current, previous int) bool {
	difference := current - previous
	fmt.Println("difference::", difference, current, previous)
	if difference > 0 {
		return direction < 0
	}
	return direction > 0
}
