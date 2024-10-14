package main

import "fmt"

func main() {
	arr := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
	val := LargestRange(arr)
	fmt.Println(val)
}



func LargestRange(array []int) []int {
	// Write your code here.
	res := [2]int{}
	visitedMap := make(map[int]bool)
	maxLengthSubArray := 0

	for i := 0; i < len(array); i ++ {
		if _, found := visitedMap[array[i]]; !found {
			visitedMap[array[i]] = true
		}
	}
	fmt.Println("visitedMap -> ", visitedMap)

	for _, num := range array {
		// check if current num-1 exists in visitedMap
		if !visitedMap[num-1]{
			// its is start of new range
			currentLength := 1
			currentNum := num

			// expand the range by checking subsequent consective number
			for visitedMap[currentNum + 1]{
				currentNum++
				currentLength++
			}

			if currentLength > maxLengthSubArray {
				maxLengthSubArray = currentLength
				res[0] = num
				res[1] = currentNum
			}
		}
		fmt.Println(num)
	}
	return []int{res[0], res[1]}
}