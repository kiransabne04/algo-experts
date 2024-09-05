package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("aslfl;nlnlflkn")
	words := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	val := GroupAnagrams(words)
	fmt.Println(val)
}

func GroupAnagrams(words []string) [][]string {
	// Write your code here.
	wordsMap := make(map[string][]string)
	for i, v := range words {
		fmt.Println(i, v)
		strRune := []rune(v)
		sort.Slice(strRune, func(i, j int) bool {
			return strRune[i] < strRune[j]
		})
		fmt.Println(string(strRune))
		if w, ok := wordsMap[string(strRune)]; ok {
			wordsMap[string(strRune)] = append(w, v)
		} else {
			wordsMap[string(strRune)] = append(w, v)
		}
		
	}
	fmt.Println(wordsMap)
	return nil
}