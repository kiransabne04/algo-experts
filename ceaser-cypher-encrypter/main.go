package main

import (
	"fmt"
	"strings"
)


func main() {
	newStr := CaesarCipherEncryptor("xyz", 2)
	fmt.Println(newStr)
}

func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
    newLetter := make([]string, 0)
    newKey := key % 26
    for _, v := range(str){
        newLetter = append(newLetter, getNewLetter(v, newKey))
    }
	fmt.Println(newLetter)
	return strings.Join(newLetter, "")
}

func getNewLetter(letterCode rune, num int) string{
	
	newLetterCode := int(letterCode) + num
	fmt.Println("letterCode -> ", letterCode, "num -> ", num, newLetterCode)
	if newLetterCode <= 122 {
		return string(newLetterCode)
	} else {
		return string(96 + int(newLetterCode) % 122)
	}
}


