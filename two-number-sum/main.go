package main

import "fmt"

func main(){
	arr := []int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}
	val := TwoNumberSum(arr, 164)
	fmt.Println(val)
}

func TwoNumberSum(array []int, target int) []int {
	var res []int
	seenMap := make(map[int]bool)
	for _, v := range array {
		fmt.Println(v)
		if _, found := seenMap[target-v]; found {
			res = append(res, v, target-v)
		} else {
			seenMap[v]=true
		}
	}
	return res
}
