package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello - generate valid ip addresses")
	inputStr := "19216802"

	results := validIPAddress(inputStr)
	fmt.Println(results)
}

func validIPAddress(input string) []string {
	fmt.Println("input -< ", input)
	result := make([]string, 0)
	if len(input) < 4 || len(input) > 12 {
		return result
	}
	backtracking(input, 0, []string{}, &result)
	return result
}

func backtracking(inputStr string, start int, parts []string, result *[]string) {
	fmt.Println("first ", inputStr, start, parts, result)
	if len(parts) == 4 {
		if len(inputStr) == start {
			*result = append(*result, fmt.Sprintf("%s.%s.%s.%s", parts[0], parts[1], parts[2], parts[3]))
		}
		return
	}

	for i := 1; i <= 3; i++ {
		if start+i > len(inputStr) {
			break
		}
		part := inputStr[start : start+i]

		if isPartValid(part) {
			//fmt.Println(part)
			backtracking(inputStr, start+i, append(parts, part), result)
		}
	}
	fmt.Println("end ", inputStr, start, parts, result)
}

func isPartValid(str string) bool {
	if len(str) == 0 || len(str) > 3 {
		return false
	}

	if len(str) > 1 && str[0] == '0' {
		return false
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}

func isValidPart1(inputPart string) bool {
	if len(inputPart) == 0 || len(inputPart) > 3 {
		return false
	}

	if len(inputPart) > 1 && inputPart[0] == '0' {
		return false
	}

	num, err := strconv.Atoi(inputPart)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}
