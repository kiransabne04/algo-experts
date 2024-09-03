package main

import (
	"fmt"
	"strings"
)

func main(){
	words := []string{"diaper", "abc", "test", "cba", "repaid"}
	val := Semordnilap(words)
	fmt.Println(val)
}

func Semordnilap(words []string) [][]string {
	// Write your code here.
    wordMap := make(map[string]int)
	pairs := [][]string{}

    for _, v := range words {
        wordMap[v]++
    }
	fmt.Println(wordMap)

	
	for _, v := range words {
		
       reverse := reverseStr(v)
	   fmt.Println("reverse -> ", reverse)
	   if wordMap[reverse] != 0 && v != reverse {
		pairs = append(pairs, []string{v, reverse})	
		delete(wordMap, v)
		delete(wordMap, reverse)
	   }
	   
    }
	
	return pairs
}

func reverseStr (word string) string {
	reverse := make([]string, 0) 
	for i := len(word)-1; i >= 0; i--{
		//fmt.Println(string(word[i]))
		reverse = append(reverse, string(word[i]))
	}
	return strings.Join(reverse, "")
}