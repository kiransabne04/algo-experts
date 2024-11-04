package main

import "fmt"

// Problem: Given a linked list, reorder it to follow the pattern L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2 ....
// Approach: Find the middle of the list, reverse the second half, and merge the two halves.
// Input: 1 -> 2 -> 3 -> 4 -> 5 -> null
// Expected Output: 1 -> 5 -> 2 -> 4 -> 3 -> null

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func reorderLinkedList(head *LinkedList) *LinkedList {
	//find the middle node
	slowPtr, fastPtr := head, head
	var previousNode *LinkedList
	for fastPtr != nil && fastPtr.Next != nil {
		previousNode = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	fmt.Println("slowPtr & fastPtr ", slowPtr, fastPtr)

	previousNode.Next = nil
	secondPart := reverseLinkedList(slowPtr)

	displayNodes(head)
	displayNodes(secondPart)

	return nil
}

func reverseLinkedList(head *LinkedList) *LinkedList {
	var previousNode, nextNode *LinkedList
	currentNode := head
	// 3 -> 4 -> 5 -> nil
	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	return previousNode
}

func displayNodes(list *LinkedList) {
	for list != nil {
		fmt.Printf("%d -> ", list.Value)
		list = list.Next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 2}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next = &LinkedList{Value: 5}

	val := reorderLinkedList(list)

	fmt.Println("val -> ", val)
}
