package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {
	//list := LinkedList{}
	// Create a linked list with duplicates for testing
	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 1}
	list.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next.Next = &LinkedList{Value: 3}

	val := RemoveDuplicatesFromLinkedList(list)
	fmt.Println("val -> ", val)
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {

	// Write your code here.
	if linkedList == nil || linkedList.Next == nil {
		return linkedList
	}

	currentNode := linkedList

	for currentNode!= nil && currentNode.Next != nil {
		if currentNode.Next.Value == currentNode.Value {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	fmt.Println(linkedList)
	return linkedList
}

// Display prints all the elements of the linked list.
func Display(list *LinkedList) {
	current := list
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}