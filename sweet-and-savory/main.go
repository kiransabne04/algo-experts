package main

import (
	"fmt"
	"math"
	"sort"
)

func main(){
	dishes := []int{2, 5, -4, -7, 12, 100, -25}
	target := -7
	val := SweetAndSavory(dishes, target)
	fmt.Println(val)
}

func SweetAndSavory(dishes []int, target int) []int {
	// Write your code here.
	res := []int{0,0}
	sort.Ints(dishes)

	fmt.Println("dishes -> ", dishes)
	start := 0
	end := len(dishes) - 1
	bestDiff := math.MaxInt64

	for start < end && dishes[start] < 0 && dishes[end] > 0{
		fmt.Println("start & end -> ", start, end)
		currentSum := dishes[start] + dishes[end]
		fmt.Println("currentSum -> ", currentSum)

		if currentSum <= target {
			currentDiff := target - currentSum
			fmt.Println("currentDiff -> ", currentDiff)
			if currentDiff < bestDiff {
				bestDiff = currentDiff
				res[0] = dishes[start]
				res[1] = dishes[end]
			}
			fmt.Println("currentDiff , bestDiff-> ", currentDiff, bestDiff, res)
			start++
		} else {
			end --
		}
	}

	return res
}
