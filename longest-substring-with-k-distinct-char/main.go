package main

import "fmt"

// Longest Substring with K Distinct Characters
// Problem: Given a string, find the length of the longest substring that contains no more than K distinct characters.
// Approach: Sliding window with character frequency tracking.
// Example:
// plaintext
// Copy code
// Input: s = "araaci", K = 2
// Output: 4
// Explanation: The longest substring with 2 distinct characters is "araa".

type Substring struct {
	Start int
	End   int
}

func main() {
	val := longestSubstringWithDistinctChar("araaci", 2)
	fmt.Println(val)
}

func longestSubstringWithDistinctChar(str string, k int) int {
	charIndexMap := make(map[byte]int)
	start := 0
	maxLength := 0

	for i := 0; i < len(str); i++ {
		charIndexMap[str[i]] = i //update last seen index in charIndexMap

		// if len(charIndexMap) > k then find the leftmost first element to remove from map
		if len(charIndexMap) > k {
			leftmost := i
			for _, index := range charIndexMap {
				if index < leftmost {
					leftmost = index
				}
			}

			start = leftmost + 1
			delete(charIndexMap, str[leftmost])
		}

		maxLength = max(maxLength, i-start+1)

	}
	return maxLength
}
