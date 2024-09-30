package main

import (
	"fmt"
	"strings"
)

func main() {
	val := UnderscorifySubstring("this is a test to see if it works and test", "test")
	fmt.Println(val)
}

func UnderscorifySubstring(str string, substring string) string {
	// Write your code here.
	res := make([][]int, 0)
	subLen := len(substring)
	var i int
	for i = 0; i < len(str) - subLen; i++{
		if str[i: i+subLen] == substring {
			//res = append(res, str[i : i + subLen])
			arr := make([]int, 0)
			arr = append(arr, i, i+subLen)
			res = append(res, arr)
		}
	}
	fmt.Println(res, i, len(str))
	// if any last part of the input str is remained processing
	if i < len(str) && substring == str[i:] {
		arr := make([]int, 0)
		arr = append(arr, i, i+subLen)
		res = append(res, arr)
	}

	if len(res) > 0 {
		return rebuildIndex(res, str, substring)
	} else {
		return str
	}
	
}

func rebuildIndex(arr [][]int, str, substring string) string {
	result := [][]int{arr[0]} // initialize with first element in results

	for i := 1; i < len(arr); i++ {
		lastRes := result[len(result) - 1] // gets the last element in the resultarr
		if arr[i][0] <= lastRes[1] { //if current elem first index is less than or eqyal to lastRes endIndex, then it overlaps or equals
			lastRes[1] = max(lastRes[1], arr[i][1])
		} else {
			result = append(result, arr[i]) // if its not equal or less then, then just append it in result arr
		}
	}
	resultStr := addUnderScoreAtIntervals(str, result)

	return resultStr
}

func addUnderScoreAtIntervals(s string, intervals [][]int) string {
	fmt.Println("intervals: ", intervals)
	var builder strings.Builder
	currentIdx := 0

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		// write the part of str before the interval into builder
		builder.WriteString(s[currentIdx:start])

		// add the underscore at the start of the interval
		builder.WriteString("_")

		// add the interval substrimg part of str for the end
		builder.WriteString(s[start : end])
		//fmt.Println(s[start : end])
		// add the underscpre at the end of the interval substirng
		builder.WriteString("_")
		fmt.Println(currentIdx, builder.String())
		//update the current index to the interval substring end, so as to continue for rest of intervals
		currentIdx = end 
	}
	// add remaining part of the stirng if left any after the last interval end point also
	if currentIdx < len(s) {
		builder.WriteString(s[currentIdx:])
	}
	return builder.String()
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}