package main

import "fmt"

// Problem: Given a string, find the length of the longest substring without repeating characters.
// Approach: Sliding window with a set or map to track characters.
// Example:
// plaintext
// Copy code
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with a length of 3.
func main() {
	val := longestSubstringWithoutRepeatingChar("abcabcbb")
	fmt.Println(val)
}

func longestSubstringWithoutRepeatingChar(str string) int {
	start := 0
	maxLength := 0
	charIndexMap := make(map[rune]int)

	for i, char := range str {
		fmt.Println("i, char ", i, char)
		// check char if exist in map, if exists retreive the previous existing index & increment, since it shows repetatio

		if val, found := charIndexMap[char]; found {

			if start <= val+1 {
				start = val + 1
				fmt.Println("start < val ", i, string(char), start, val, found, val+1)
			}
		}

		currentLength := i - start + 1
		maxLength = max(maxLength, currentLength)

		// update index map to lastest index
		charIndexMap[char] = i
	}

	fmt.Println("charIndexMap -> ", charIndexMap)
	return maxLength
}
