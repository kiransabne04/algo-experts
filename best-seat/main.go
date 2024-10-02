package main

import "fmt"

func main(){
	array := []int{1,0,1,0,0,0,1}
	res := BestSeat(array)
	fmt.Println(res)
}

func BestSeat(seats []int) int {
	// Write your code here.
	res := -1

	// Check if the window size is valid
	if len(seats) < 3 {
		fmt.Println("Window size is larger than array size")
		return res
	}

	

	return res
}
