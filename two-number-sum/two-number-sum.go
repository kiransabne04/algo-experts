package main

// two numbers from the given array should be equal to targetSum
import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello World")
	ans := TwoNumberSum3([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	fmt.Println(ans)
}

func TwoNumberSum1(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			//fmt.Println(array[i])
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}

func TwoNumberSum2(array []int, target int) []int {
	m := make(map[int]bool, 0)
	for i := 0; i < len(array); i++ {
		_, val := m[target-array[i]]
		fmt.Println(val)
		if val {
			return []int{array[i], target - array[i]}
		} else {
			m[array[i]] = true
		}
		fmt.Println("m -> ", m)
	}
	return []int{}
}

func TwoNumberSum3(array []int, target int) []int {
	sort.Ints(array)
	fmt.Println("array -> ", array)

	firstPointer := 0
	lastPointer := len(array) - 1
	for firstPointer < lastPointer {
		sum := array[firstPointer] + array[lastPointer]
		if sum == target {
			return []int{array[firstPointer], array[lastPointer]}
		} else if sum < target {
			firstPointer++
		} else if sum > target {
			lastPointer--
		}
	}

	return []int{}
}
