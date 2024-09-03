package main

import (
	"fmt"
)

func main() {
	val := LongestPalindromicSubstring("abaxyzzyxf")
	fmt.Println(val)
}

func LongestPalindromicSubstring(str string) string {
	if len(str) == 0 {
		return ""
	}
	longest := ""

	for i := 0; i < len(str); i++{
		// check for odd length where center is a letter
		oddPalindrome := expandAroundCenter(str, i, i)
		if len(oddPalindrome) > len(longest) {
			longest = oddPalindrome
		}
		// center is between letter
		evenPalindrome := expandAroundCenter(str, i, i+1)
		if len(evenPalindrome) > len(longest) {
			longest = evenPalindrome
		}
	}
	return longest
} 

func expandAroundCenter(str string, leftIndx, rightIndx int) string {
	for leftIndx >= 0 && rightIndx < len(str) && str[leftIndx] == str[rightIndx]{
		leftIndx--
		rightIndx++
	}
	return str[leftIndx+1 : rightIndx]
}