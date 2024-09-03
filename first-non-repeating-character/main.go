package main

import "fmt"

func main (){
	val := FirstNonRepeatingCharacter("abcdcaf")
	fmt.Println(val)
}

func FirstNonRepeatingCharacter(str string) int {
	// Write your code here.
	charCount := make(map[rune]int)
	for _, v := range str {
		charCount[v]++
	}

	for i, v := range str {
		if charCount[v] == 1 {
			return i
		}
	}
	fmt.Println("charCount -> ", charCount)
	return -1
}
