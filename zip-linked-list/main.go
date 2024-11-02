package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ZipLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	smallPtr := linkedList
	fastPtr := linkedList

	var previousNode *LinkedList
	for fastPtr != nil && fastPtr.Next != nil {
		previousNode = smallPtr
		smallPtr = smallPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	fmt.Println("smallPtr ", smallPtr, "previousNode -> ", previousNode," fastPtr ", fastPtr)

	// if fastPtr != nil {
	// 	smallPtr = smallPtr.Next
	// }
	previousNode.Next = nil

	firstPart := linkedList
	secondPart := reverseLinkedList(smallPtr)
	displayNodes(firstPart)
	displayNodes(secondPart)

	for firstPart != nil && secondPart != nil {
		// temp poninters to next nodes
		firstNext := firstPart.Next
		secondNext := secondPart.Next

		//zipping nodes togther
		firstPart.Next = secondPart
		if firstNext == nil {
			break
		}
		secondPart.Next = firstNext

		firstPart = firstNext
		secondPart = secondNext

	}
	displayNodes(linkedList)
	return linkedList
}

func reverseLinkedList(list *LinkedList) *LinkedList {
	var previousNode *LinkedList
	current := list
	var nextNode *LinkedList
	for current != nil {
		nextNode = current.Next
		current.Next = previousNode
		previousNode = current
		current = nextNode
	}
	return previousNode
}

func displayNodes (list *LinkedList) {
	for list != nil {
		fmt.Printf(" %d -> ", list.Value)
		list = list.Next
	}
	
	fmt.Println()
}

func main(){
	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 2}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next = &LinkedList{Value: 5}
	list.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	list.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 10}

	val := ZipLinkedList(list)
	fmt.Println("val -> ",val)
}