package main

import (
	"fmt"
	"strconv"
	"strings"
)

//"aaaaaaaaaaaa" => "9a3a"
// "aaaaaaaaaaaaabbccccdd" => "9a4a2b4c2d"

func main(){
	val := RunLengthEncoding("aaaaaaaaaaaa")
	
	fmt.Println(val)
}

func RunLengthEncoding(str string) string {
	res := make([]string, 0)
	count := 1
	previous := string(str[0])
	fmt.Println(" count -> ", count, previous)
	for i := 1; i < len(str); i++{
		fmt.Println(i, str[i])
		if(count < 9) && (previous == string(str[i])){
			count ++
			previous = string(str[i])
		} else {
			res = append(res, strconv.Itoa(count), previous)
			count = 1
			previous = string(str[i])
		}
		fmt.Println("b count -> ", count, previous)
	}
	res = append(res, strconv.Itoa(count), previous)
	return strings.Join(res, "")
}