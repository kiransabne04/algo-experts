package main

import (
	"fmt"
)

func main() {
	input := []string{"this", "that", "did", "deed", "them!", "a"}
	val := MinimumCharactersForWords(input)
	fmt.Println(val)
}

func MinimumCharactersForWords(words []string) []string {
	// Write your code here.
	globalCharacterCount := make(map[rune]int)

	for i, word := range words {
		fmt.Println(word, i)
		wordCharacterFrequency := wordCharacterCount(word)
		updateGlobalFrequencyMap(globalCharacterCount, wordCharacterFrequency)
		fmt.Println("wordCharacterF -> ", wordCharacterFrequency, globalCharacterCount)

	}
	//
	
	return makeSliceOfCharacterFrequency(globalCharacterCount)
	//t, h, i, s, a, t, d, d, e, e, m, !

}

func makeSliceOfCharacterFrequency(globalCharacterCount map[rune]int) []string {
	characterArr := make([]string, 0)
	for key, val := range globalCharacterCount {
		for i := 0 ; i < val; i++ {
			characterArr = append(characterArr, string(key))
		}
	}
	return characterArr
}

func wordCharacterCount(str string) map[rune]int{
	localCharacterFreqMap := make(map[rune]int)
	for _, character := range str {
		localCharacterFreqMap[character]++
	}
	return localCharacterFreqMap
}

func updateGlobalFrequencyMap(globalCharacterCount map[rune]int, wordCharacterFrequency map[rune]int) {
	for character, count := range wordCharacterFrequency {
		if characterCount, found := globalCharacterCount[character]; found {
			globalCharacterCount[character] = max(count, characterCount)
		} else {
			globalCharacterCount[character] = count
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



