package main

import (
	"fmt"
)

func main() {
	val := ValidIPAddresses("1921680")
	fmt.Println(val)
}

func ValidIPAddresses(str string) []string {
	// Write your code here.
	// lets create a list of IP sequences and then we will check for its validity
	start := 1
	//part := make([]string, 0)
	for i := 0; i < 3; i++ {
		// if the start+i position is > len of inputStr
		if start+i > len(str) {
			break
		}

		fmt.Println(string(str[i]), start)

		//part = append(part, string(str[i]))
		part := str[start : start+i]
		//start = start + i
		fmt.Println("startPntr -> ", string(part))
	}
	return []string{}
}
