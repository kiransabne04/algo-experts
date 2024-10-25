package main

import "fmt"

// The idea is to first find middle of a linked list using two pointers, first one moves one at a time and second one moves two at a time. When second pointer reaches end, first reaches middle. We also keep track of previous of first pointer so that we can remove middle node from its current position and can make it head.
// Input  : 1 2 3 4 5
// Output : 3 1 2 4 5

// Input  : 1 2 3 4 5 6
// Output : 4 1 2 3 5 6
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {
	// Create a linked list with duplicates for testing
	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 1}
	list.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next.Next = &LinkedList{Value: 3}
	val := setMiddleNodeHead(list)
	fmt.Println("val -> ", val)
	displayList(val)
}

func setMiddleNodeHead(ll *LinkedList) *LinkedList {

	currentNode, nextNode := ll, ll
	var previousNode *LinkedList
	for nextNode != nil && nextNode.Next != nil {
		//fmt.Println("currentNode.Value => ", previousNode.Value, currentNode.Value, nextNode.Value)
		previousNode = currentNode

		currentNode = currentNode.Next
		nextNode = nextNode.Next.Next

	}

	// Set the previous node's next to skip the middle node
	if previousNode != nil {
		previousNode.Next = currentNode.Next
	}
	currentNode.Next = ll
	ll = currentNode

	return ll
}

func displayList(ll *LinkedList) {
	for ll != nil {
		fmt.Printf("%d -> ", ll.Value)
		ll = ll.Next
	}
}
