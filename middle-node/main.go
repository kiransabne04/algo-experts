package main

import (
	"fmt"
)

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {

	list := &LinkedList{Value: 2}
	list.Next = &LinkedList{Value: 7}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 5}
	fmt.Println(list)
	val := MiddleNode(list)
	fmt.Println(val)
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	count := 1
	current := linkedList

	for current != nil {
		fmt.Println("current & current.Next -> ", current.Value, current.Next)
		count++
		current = current.Next
	}
	fmt.Println("count -> ", count/2)
	if count%2 == 0 {
		count = count / 2
	} else {
		count = (count / 2) + 1
	}
	fmt.Println("counrt -> ", count)
	//fmt.Println("current & current.Next -> ", current)

	for linkedList != nil {
		count--
		fmt.Println("linkedList & linkedList.Next -> ", linkedList.Value, linkedList.Next, count)
		if count == 0 {
			return linkedList
		}
		linkedList = linkedList.Next
	}

	return nil
}
