package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := ValidIPAddresses("255255255256")
	fmt.Println(val)
}

func isValidPart(val string) bool {
	if len(val) == 0 || len(val) > 3 {
		return false
	}

	// check if first character is '0', it can be '0' only if len(val) = 1 but if len(val) > 1 then first character should not be 0
	if len(val) > 1 && val[0] == '0' {
		return false
	}

	//convert string val to int and check the range between 0 & 255
	input, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return input >= 0 && input <= 255
}

func ValidIPAddresses(str string) []string {
	// Write your code here.
	var result []string
	if len(str) < 4 || len(str) > 12 {
		return result
	}
	backtracking(str, 0, []string{}, &result)

	return result
}

func backtracking(str string, start int, parts []string, result *[]string) {
	//start := 0
	if len(parts) == 4 {
		if len(str) == start {
			*result = append(*result, fmt.Sprintf("%s.%s.%s.%s", parts[0], parts[1], parts[2], parts[3]))
		}
		return
	}

	for startPtr := 1; startPtr <= 3; startPtr++ {
		if start+startPtr > len(str) {
			break
		}
		part := str[start : start+startPtr]

		if isValidPart(part) {
			backtracking(str, start+startPtr, append(parts, part), result)
			fmt.Println("part => ", part)
		}

	}
}
