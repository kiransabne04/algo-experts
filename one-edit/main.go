package main

import "fmt"

func main() {
	stringOne := "abcdefghijklmnopqrstuvwxyz"
	stringTwo := "abcdefghijklmnopqrstuvwxyz"
	res := OneEdit(stringOne, stringTwo)
	fmt.Println(res)
}

func OneEdit(stringOne string, stringTwo string) bool {
	// Write your code here.
	source, target := stringOne, stringTwo
	
	if len(source) < len(target) {
		source, target = target, source
	}
	if len(source) - len(target) > 1 {
		return false
	}

	for i := 0; i < len(target); i++ {
		fmt.Println(string(target[i]), " ", string(source[i]))
		if source[i] != target[i] {
			if len(source) == len(target) {
				return source[(i+1):] == target[(i+1):]
			} else {
				return source[(i+1):] == target[i:]
			}
		} 
	}
	return len(source) == len(target) || len(source) == len(target) + 1
}
