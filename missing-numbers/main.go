package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 4, 3}
	val := MissingNumbers(nums)
	fmt.Println(val)
}

func MissingNumbers(nums []int) []int {
	res := []int{}
	seenMap := make(map[int]int)
	// nums array seed the data in seenMap
	for i, v := range nums {
		seenMap[v] = i
	}

	// now loop from 1 to len(nums) + 3 .. since 2 elements will be missing and check the i exists in seenMap. it not then add in res slice declared above
	for i := 1; i < len(nums)+3; i++ {
		fmt.Println("new:: ", i)
		if _, found := seenMap[i]; !found {
			res = append(res, i)
		}
	}

	return res
}
