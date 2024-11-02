package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	// Write your code here.
	// Initialize the Starting Head:

	// Set mergedListHead to the smaller of headOne or headTwo.
	// Initialize a current pointer to keep track of the last node in the merged list as you build it.
	// Traverse Both Lists:
	
	// As you compare nodes from headOne and headTwo, set current.Next to the smaller node and move current forward to current.Next.
	// Attach Remaining Nodes:
	
	// When one list is exhausted, attach the remaining nodes of the other list to current.Next.
	// Return the Head of the Merged List:
	
	// After merging, return the original mergedListHead
	
	var mergedListHead *LinkedList
	if headOne.Value < headTwo.Value {
		mergedListHead = headOne
		headOne = headOne.Next
	} else {
		mergedListHead = headTwo
		headTwo = headTwo.Next
	}

	// current to build mergedlist
	current := mergedListHead

	fmt.Println("mergedListHead -> ", mergedListHead, "headOneList -> ", headOne, "headTwoList -> ", headTwo)
	for headOne != nil && headTwo != nil {
		fmt.Println("headOne & headTwo -=> ", headOne, headTwo)
		if headOne.Value < headTwo.Value {

			current.Next = headOne
			headOne = headOne.Next
		} else {
			current.Next = headTwo
			headTwo = headTwo.Next
		}
		current = current.Next // moving current forward
	}

	// appending remaining elements of the list to the mergedList
	if headOne != nil {
		current.Next = headOne
	} else if headTwo != nil {
		current.Next  = headTwo
	}

	return mergedListHead

}

func main() {

	headOne := &LinkedList{Value: 2}
	headOne.Next = &LinkedList{Value: 6}
	headOne.Next.Next = &LinkedList{Value: 7}
	headOne.Next.Next.Next = &LinkedList{Value: 8}

	headTwo := &LinkedList{Value: 1}
	headTwo.Next = &LinkedList{Value: 3}
	headTwo.Next.Next = &LinkedList{Value: 4}
	headTwo.Next.Next.Next = &LinkedList{Value: 5}
	headTwo.Next.Next.Next.Next = &LinkedList{Value: 9}
	headTwo.Next.Next.Next.Next.Next = &LinkedList{Value: 10}

	val := MergeLinkedLists(headOne, headTwo)
	fmt.Println("val -> ", val)

}
