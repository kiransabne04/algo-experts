package main

import "fmt"

type substring struct {
	start int
	end int
}

func main() {
	val := LongestSubstringWithoutDuplication("clementisacap")
	fmt.Println(val)
}


func LongestSubstringWithoutDuplication(str string) string {
	// Write your code here.
	seenMap := make(map[rune]int)
	lSubString := substring{0, 1}
	startIndex := 0
	for i , char := range str {
		fmt.Println("i -> ", char)
		
		if value, found  := seenMap[char]; found {
			fmt.Printf("found value %v in map & its previous position is %d\n", char, value )
			if startIndex < value + 1 {
				startIndex = value + 1
			}
		} 
		if lSubString.end - lSubString.start < i + 1 - startIndex {
			lSubString = substring{startIndex, i + 1}
		}
		seenMap[char] = i
	}
	fmt.Println("seenMap -> ", seenMap, lSubString)
	return str[lSubString.start : lSubString.end]
}