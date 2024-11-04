package main

import "fmt"

// Problem: Remove the n-th node from the end of a singly linked list in one pass.
// Approach: Use two pointers with a fixed n-node gap.
// Input: 1 -> 2 -> 3 -> 4 -> 5 -> null, n = 2
// Expected Output: 1 -> 2 -> 3 -> 5 -> null

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func removeKthNodeFromTail(list *LinkedList, k int) *LinkedList {

	currentNode, nextNode := list, list
	var previousNode *LinkedList

	i := 0 // we are starting with 0 since we want to stop at one node before the node we want to delete.

	for i < k && nextNode != nil {

		nextNode = nextNode.Next
		i++
	}

	fmt.Println("previousNode, currentNode, nextNode -> ", previousNode, currentNode, nextNode)

	if nextNode == nil {
		return list.Next
	}

	for nextNode != nil {
		previousNode = currentNode
		nextNode = nextNode.Next
		currentNode = currentNode.Next
	}

	fmt.Println("previousNode, currentNode, nextNode -> ", previousNode, currentNode, nextNode)

	if previousNode != nil {
		previousNode.Next = currentNode.Next
	}
	fmt.Println("previousNode, currentNode -> ", previousNode, currentNode)

	displayNodes(list)

	return list
}

func main() {
	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 2}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next = &LinkedList{Value: 5}

	val := removeKthNodeFromTail(list, 6)
	fmt.Println("val ", val)
}

func displayNodes(head *LinkedList) {
	for head != nil {
		fmt.Printf("%d -> ", head.Value)
		head = head.Next
	}
	fmt.Println()
}
