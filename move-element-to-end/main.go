package main

import "fmt"

func main() {
	array := []int{2, 4, 2, 5, 6, 2, 2}
	toMove := 2
	val := MoveElementToEnd(array, toMove)
	fmt.Println(val)
}

func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	leftIdx, rightIdx := 0, len(array) -1
	fmt.Println(array[leftIdx], array[rightIdx])
	for leftIdx < rightIdx {
		for leftIdx < rightIdx && array[rightIdx] == toMove {
			rightIdx--
		}
		
		if array[leftIdx] == toMove {
			array[leftIdx], array[rightIdx] = array[rightIdx], array[leftIdx]
		}
		fmt.Println(":::", leftIdx, rightIdx)
		leftIdx++
	}
	return array
}
