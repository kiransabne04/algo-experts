package main

import "fmt"

func main() {
	val := TransposeMatrix([][]int{{1,2}})
	fmt.Println(val)
}

func TransposeMatrix(matrix [][]int) [][]int {
	// Write your code here.
	transposeMatrix := make([][]int, 0)
	for i, val := range matrix[0] {
		fmt.Println("i & val ", i, val)
		newRow := make([]int, 0)
		for j, _ := range matrix {
			newRow = append(newRow, matrix[j][i])
		}
		transposeMatrix = append(transposeMatrix, newRow)
	}
	return transposeMatrix
}
