package main

import "fmt"

func main(){
	array := []int{1, 2, 3, 5, 6, 8, 9}
	val := SortedSquaredArray(array)
	fmt.Println(val)

}

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	res := make([]int, len(array))
	sIndex := 0
	eIndex := len(array)-1

	for i := len(array) - 1; i >= 0; i-- {
		smallerValue := array[sIndex]
		largerValue := array[eIndex]

		if abs(smallerValue) > abs(largerValue) {
			res[i] = smallerValue * smallerValue
			sIndex += 1
		} else {
			res[i] = largerValue * largerValue
			eIndex -= 1
		}
	}
	return res
}

func abs (a int) int {
	if a < 0 {
		return -a
	}
	return a
}
