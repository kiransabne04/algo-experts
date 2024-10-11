package main

import (
	"fmt"
	"sort"
)

func main() {
	val := FourNumberSum([]int{7, 6, 4, -1, 1, 2}, 16)
	fmt.Println(val)
}

func FourNumberSum(array []int, target int) [][]int {
	// sort the input array
	sort.Ints(array)
	quad := make([][]int, 0)
	pairsMap := make(map[int][][]int)
	return nil
}
