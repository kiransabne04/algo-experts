package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	val := LongestPeak(array)
	fmt.Println("val -> ", val)
}

func LongestPeak(array []int) int {
	// Write your code here.
	longestLength := 0
	i := 1
	for i < len(array)-1 {

		isPeak := array[i] > array[i-1] && array[i] > array[i+1]
		if !isPeak {
			i += 1
			continue
		}

		//if its peak then check direction on each side left & right +2 & -2 wrt i and see if it follows same pattern
		leftIndx := i - 2
		for leftIndx >= 0 && array[leftIndx] < array[leftIndx+1] {
			fmt.Println("l-array -> ", array[leftIndx], array[leftIndx-1])
			leftIndx -= 1
		}

		rightIndx := i + 2
		for rightIndx < len(array) && array[rightIndx] < array[rightIndx-1] {
			fmt.Println("r-array -> ", array[rightIndx], array[rightIndx-1])
			rightIndx += 1
		}

		fmt.Println("Indx:", leftIndx, rightIndx)
		currentPeakLength := rightIndx - leftIndx - 1
		if currentPeakLength > longestLength {
			longestLength = currentPeakLength
		}
		i = rightIndx
	}

	return longestLength
}
