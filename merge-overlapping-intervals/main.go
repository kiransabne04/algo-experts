package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1,2}, {3,5}, {4,7}, {6,8}, {9,10}} // {1,2}, {3,8}, {9,10}
	res := MergeOverlappingIntervals(intervals)
	fmt.Println(res)
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	// Write your code here.

	//sort intervals by their starting value
	sort.Slice(intervals,func(i, j int) bool {return intervals[i][0] < intervals[j][0]})

	res := [][]int{}
	res = append(res, intervals[0])

	// Iterate over each interval starting from the second one
	for i := 1; i < len(intervals); i++ {
		
		interval := intervals[i]
		prevInterval := res[len(res)-1] // Get the last interval added to the result
		
		// Check if the current interval overlaps with the previous one
		if prevInterval[1] >= interval[0] {
			
			res[len(res)-1][1] = max(prevInterval[1], interval[1])
			
		} else {
			// No overlap, append the current interval to the result
			res = append(res, interval)
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
