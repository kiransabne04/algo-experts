package main

import (
	"fmt"
)

func main() {
	val := FourNumberSum([]int{7, 6, 4, -1, 1, 2}, 16)
	fmt.Println(val)
}

func FourNumberSum(array []int, target int) [][]int {
	// sort the input array
	
	quad := [][]int{}
	pairsMap := make(map[int][][]int)

	for i := 1; i < len(array) - 1; i ++{
		
		for j := i + 1; j < len(array); j ++{
			fmt.Println(array[i], array[j])
			diff := target - (array[i] + array[j])
			if pairs, found := pairsMap[diff]; found {
				// pairs can have multiple array
				for _, pair := range pairs {
					quad = append(quad, []int{pair[0], pair[1], array[i], array[j]})
				}
				fmt.Println(pairs)
			}
		}

		// adding pairs for the previous and current of i
		for k := 0; k < i; k++ {
			fmt.Println("k & i array val -> ", array[k], array[i])
			currentSum := array[k] + array[i]
			pairsMap[currentSum] = append(pairsMap[currentSum], []int{array[k], array[i]})
		}
	}

	fmt.Println(pairsMap)
	return quad
}
