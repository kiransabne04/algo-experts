package main

import "fmt"

//Problem: Merge two sorted linked lists and return a single sorted list.
// List 1: 1 -> 3 -> 5 -> null
// List 2: 2 -> 4 -> 6 -> null
// Expected Output: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null

type LinkedList struct {
	Value int
	Next *LinkedList
}

func mergeLinkedList(list1, list2 *LinkedList) *LinkedList {

	var head *LinkedList
	if list1.Value < list2.Value {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	current := head

	displayNodes(current)

	for list1 != nil && list2 != nil {
		if list1.Value < list2.Value {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	displayNodes(head)

	return nil
}

func displayNodes (head *LinkedList) {
	for head != nil {
		fmt.Printf(" %d-> ", head.Value)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	list1 := &LinkedList{Value: 1}
	list1.Next = &LinkedList{Value: 3}
	list1.Next.Next = &LinkedList{Value: 5}

	list2 := &LinkedList{Value: 2}
	list2.Next = &LinkedList{Value: 4}
	list2.Next.Next = &LinkedList{Value: 6}

	_ = mergeLinkedList(list1, list2)
}