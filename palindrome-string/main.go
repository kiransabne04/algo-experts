package main

import "fmt"

func main(){

	val := isPalindrome("abcdcba", 0, len("abcdcba") - 1)
	fmt.Println(val)
}

func isPalindrome(str string, start, end int) bool {
	if start >= end {
		return true
	}
	if str[start] == str[end] {
		return isPalindrome(str, start + 1, end - 1)
	}
	return false
}