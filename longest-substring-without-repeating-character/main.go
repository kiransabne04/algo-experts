package main

import "fmt"

func main() {
	val := longestSubstringWithoutRepeat("eghghhgg")
	fmt.Println(val)
}

func longestSubstringWithoutRepeat(str string) string {
	left := 0
	seenMap := make(map[rune]int)
	maxStr := ""
	maxLen := 0

	for idx, char := range str {

		// for char present in the map & val is in current window
		if val, found := seenMap[char]; found && val >= left {
			left = val + 1
		}

		seenMap[char] = idx
		// check if curent window is longest
		if idx-left+1 > maxLen {
			maxLen = idx - left + 1
			maxStr = str[left : idx+1]
		}
	}
	return maxStr
}
