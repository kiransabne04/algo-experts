package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	fmt.Println("Hello *** ")
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	// var a, b int
	// scanf("%d %d\n", &a, &b)
	// printf("%d\n", a+b)
	var n int
	scanf("%d\n", &n)
	arr := make([]int, 0)
	var t int
	for i := 0; i < n; i++ {
		scanf("%d\n", &t)
		arr = append(arr, t)
	}
	printf("%d\n", arr)
	printf("ans is -> %d\n", NonConstructibleChange(arr))
}

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	currentChange := 0
	for _, coin := range coins {
		if coin > currentChange+1 {
			return currentChange + 1
		}

		currentChange += coin
	}
	return currentChange + 1
}
