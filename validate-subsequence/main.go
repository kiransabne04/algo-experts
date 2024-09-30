package main

import "fmt"

func main(){
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{22, 25, 6}
	val := IsValidSubsequence(array, sequence)
	fmt.Println(val)
}


func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	//res := false
	sIndex := 0
	for _, val := range array {
		fmt.Println(val)
		if val == sequence[sIndex] {
			sIndex++
		}
		fmt.Println("sIndex -> ", sIndex)

		if sIndex == len(sequence) {
			break
		}
	}
	return sIndex == len(sequence)
}
