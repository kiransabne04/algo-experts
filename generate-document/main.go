package main

import "fmt"

func main(){
	val := GenerateDocument("abcabc", "")
	fmt.Println(val)
}

func GenerateDocument(characters string, document string) bool {
	// Write your code here.
	fmt.Println(characters, document)
    if document == "" {
        return true
    }

    characterMap := make(map[rune]int)
    for _, v := range characters {
        characterMap[v]++
    }

	// checking in provided document
	for _, v := range document {
		characterMap[v]--
		if characterMap[v] < 0 {
			return false
		}
	}
    
	fmt.Println(characterMap)
	return true
}