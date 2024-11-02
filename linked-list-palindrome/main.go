package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func LinkedListPalindrome(head *LinkedList) bool {
	// Write your code here.
	// find the middle node of the list

	if head.Next == nil {
        return true // since its a single node, then its valid palindrone
    }

	slowPtr := head
	fastPtr := head

	var previousNode *LinkedList
	for fastPtr != nil && fastPtr.Next != nil {
		//fmt.Println("slowPtr -> ", slowPtr.Value, "fastPtr -> ", previousNode)
		previousNode = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	fmt.Println("slowPtr -> ", slowPtr.Value, "fastPtr -> ", fastPtr, "previousNode -> ", previousNode)

	// if fastPtr.Val is nil then its even length list else its odd length list
	if fastPtr != nil {
		//previousNode = slowPtr
		slowPtr = slowPtr.Next
	}
	// Split the list into two parts
	previousNode.Next = nil // End the first half
	secondPart := slowPtr   // Start of the second half

	displayNodes(head)
	displayNodes(secondPart)
	
	// reverse the secondPart linked list
	reversedList := reverseLinkedList(secondPart)
	displayNodes(reversedList)

	// now check if both the node values are same
	firstPart := head
	for firstPart != nil && reversedList != nil {
		if firstPart.Value != reversedList.Value {
			return false
		}
		firstPart = firstPart.Next
		reversedList = reversedList.Next
	}

	displayNodes(head)

	return true
}

func reverseLinkedList(list *LinkedList) *LinkedList {

	var previousNode *LinkedList
	current := list
	var nextNode *LinkedList

	for current != nil {
		// 2 -> 1 -> 0 -> nil
		// 0 -> 1 -> 2 -> nil
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
	list := &LinkedList{Value: 0}
	list.Next = &LinkedList{Value: 1}
	list.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next.Next = &LinkedList{Value: 1}
	list.Next.Next.Next.Next.Next = &LinkedList{Value: 0}
	//list.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 10}

	val := LinkedListPalindrome(list)
	fmt.Println("val -> ",val)
}