package main

import (
	"fmt"
	s "strings"
)

func main(){
	val := CommonCharacters([]string{"abc", "bcd", "cbaccd"})
	fmt.Println(val)
}

func CommonCharacters(strings []string) []string {
	// Write your code here.
	smallestStr := getSmallestString(strings)
	fmt.Println("smallestStr -> ", smallestStr)

	// seed map from smallestStr
	charCount := make(map[string]int)
	for _, v := range smallestStr {
		charCount[string(v)]++
	}
	fmt.Println(charCount)
	
	// eliminate chars frm maps if they are not in strings
	for i := 0; i < len(strings); i++ {
		processStringAndCharacter(strings[i], charCount)
	}
	//fmt.Println(charCount)

	// make a slice of required length
	res := make([]string, len(charCount))
	i := 0
	for k := range charCount {
		res[i] = k
		i++
	}

	return res
}

func processStringAndCharacter(str string, charCount map[string]int) {
	for k, _ := range charCount {
		if !s.Contains(str, k) {
			delete(charCount, k)	
		}
	}
}

func getSmallestString(arr []string) string {
	smallestString := arr[0]
	for _, v := range arr {
		if len(smallestString) > len(v)  {
			smallestString = v
		}
	}
	return smallestString
}


