package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := ValidIPAddresses("1921680")
	fmt.Println(val)
}

func ValidIPAddresses(str string) []string {
	// Write your code here.
	ipNum, err := strconv.Atoi(str)
	if err != nil {
		
	}
	return []string{}
}