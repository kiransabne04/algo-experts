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

	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 2}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next = &LinkedList{Value: 5}
	list.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	list.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 7}
	list.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 8}

	fmt.Println(list)
	val := MiddleNode1(list)
	fmt.Println(val)
}

func MiddleNode1(ll *LinkedList) *LinkedList {
	slowPtr := ll
	fastPtr := ll

	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
		if fastPtr != nil { // Check if fastPtr is not nil before accessing its value
			fmt.Println("SlowPtr, fastPtr -> ", slowPtr.Value, fastPtr.Value)
		}
	}

	return slowPtr
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
