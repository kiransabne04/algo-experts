package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}))
}

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	sIndex := 0
	for i := 0; i < len(array); i++ {
		if sIndex == len(sequence) {
			break
		}
		if sequence[sIndex] == array[i] {
			sIndex++
		}
	}
	return sIndex == len(sequence)
}
