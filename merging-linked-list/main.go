package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergingLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	seenMap := make(map[int]bool)

	for linkedListOne != nil {
		seenMap[linkedListOne.Value] = true
		linkedListOne = linkedListOne.Next
	}
	
	currentNode2 := linkedListTwo
	for currentNode2 != nil {
		if _, found := seenMap[currentNode2.Value]; found {
			return currentNode2
		}
		currentNode2 = currentNode2.Next
	}
	return nil
}

func MergingLinkedLists1(ll1 *LinkedList, ll2 *LinkedList) *LinkedList {
	curr1, curr2 := ll1, ll2

	for curr1 != curr2 {
		if curr1 == nil {
			curr1 = ll2
		} else {
			curr1 = curr1.Next
		}

		if curr2 == nil {
			curr2 = ll1
			curr2 = curr2.Next
		}
	}
	return curr1
}