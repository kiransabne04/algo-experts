package main

import (
	"fmt"
	"math"
)

func main (){
	array := []int{2}
	res := MajorityElement(array)
	fmt.Println("res => ", res)
}

func MajorityElement(array []int) int {
	// Write your code here.
	seenMap := make(map[int]int)
	for _, v := range array {
		seenMap[v]++
	}

	fmt.Println(seenMap)
	maxValue := math.MinInt32
	majorityElem := 0
	for k, v := range seenMap {
		if v > maxValue {
			maxValue = v
			majorityElem = k
		}
	}

	return majorityElem
}

//using boyer-moore majority vote alogrith, for finding the majority in a sequence of elements using linear time & constant numer of space memory
func MajorityElement1(arr []int) int {
	currentNum := arr[0]
	currentFreq := 1

	for i := 1; i < len(arr); i++{
		if arr[i] == currentNum {
			currentFreq ++;
		} else {
			currentFreq --;
		}

		if currentFreq < 0 {
			currentNum = arr[i]
			currentFreq = 1
		}
	}
	return currentNum

}