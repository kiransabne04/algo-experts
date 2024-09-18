package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("aslfl;nlnlflkn")
	words := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	val := GroupAnagrams(words)
	fmt.Println(val)
}

func GroupAnagrams(words []string) [][]string {
	wordsMap := make(map[string][]string)
	for i, v := range words {

		wordStr := []rune(v)

		sort.Slice(wordStr, func(i, j int) bool {
			return wordStr[i] < wordStr[j]
		})

		fmt.Println("v ", i, v, string(wordStr))
		if w, ok := wordsMap[string(wordStr)]; ok {
			wordsMap[string(wordStr)] = append(w, v)
		} else {
			wordsMap[string(wordStr)] = append(w, v)
		}
	}

	fmt.Println(wordsMap)
	wordList := make([][]string, 0)
	for _, v := range wordsMap {
		wordList = append(wordList, v)
	}
	return wordList
}
