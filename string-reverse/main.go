package main

import "fmt"

func main(){
	fmt.Printf("%v\n", reverse("abcdef"))
}

func reverse(s string) string {
	chars := []rune(s)
	fmt.Println(chars)
	for i, j := 0, len(s)-1 ; i < j; i, j = i + 1, j - 1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

