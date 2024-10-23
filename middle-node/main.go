package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {

	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 1}
	list.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next.Next = &LinkedList{Value: 3}

	val := MiddleNode(list)
	fmt.Println(val)
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	count := 1
	current := linkedList

	for linkedList.Next != nil {

	}

	return nil
}
