package main

import "fmt"

func main() {
	val := generateSubset([]int{1, 2, 3})
	fmt.Println("val -> ", val)
}

func generateSubset(nums []int) [][]int {
	var result [][]int
	backtracking(nums, 0, []int{}, &result)
	return result
}

func backtracking(nums []int, index int, current []int, result *[][]int) {
	//fmt.Println("nums -> ", nums)
	temp := make([]int, len(current))
	copy(temp, current)
	fmt.Println("nums -> ", nums)
	fmt.Println("temp -> ", temp)
	*result = append(*result, temp)
	fmt.Println("result -> ", result)
	for i := index; i < len(nums); i++ {
		current = append(current, nums[i])

		backtracking(nums, i+1, current, result)
		current = current[:len(current)-1]
		fmt.Println("current -> ", current)
	}

}
