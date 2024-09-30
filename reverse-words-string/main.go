package main

import (
	"fmt"
	"unicode"
)

func main() {
	inputStr := "whitespaces    4"
	fmt.Println(reverseWordStr(inputStr))
}

func reverseWordStr(str string) string {

	// to solve, we can reverse the whole sentence first and the reverse the words in each sentence
	runes := []rune(str)
	// reverse the sentence
	reverse(runes, 0, len(runes)-1)
	fmt.Println("runes -> ", runes)

	// now search start and end letter form the words in the string
	start := 0
	for start < len(runes){
		for start < len(runes) && unicode.IsSpace(runes[start]) {
			start ++ // if its space then ignore and increment start point till finding a char
		}

		end := start // after finding the first non-space character for the start pointer, we need to find the end non-space pointer. 
		// to find ending non-space character we will start from start pointer and look for the non-space character
		for end < len(runes) && !unicode.IsSpace(runes[end]) {
			end++
		}

		// get characters in between start and end and reverse it
		if (start < end) {
			reverse(runes, start, end - 1)
		}

		// after reversing the runes part, continue to move the start pointer to end pointer, so as to move forward with remaining parts
		start = end
	}
	return string(runes)
}

func reverse(chars []rune, start, end int) {
	fmt.Println("reverse chars recevied ", chars)
	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
}