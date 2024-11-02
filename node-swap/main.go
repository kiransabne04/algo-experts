package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func NodeSwap(head *LinkedList) *LinkedList {
	// Write your code here.

	dummyNode := &LinkedList{Next: head}
	previousNode := dummyNode
	
	for previousNode.Next != nil && previousNode.Next.Next != nil {
		firstNode := previousNode.Next
		secondNode := firstNode.Next

		//swap here 3 nodes
		previousNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		previousNode = firstNode
 	}
	displayNodes(dummyNode.Next)
	return dummyNode.Next
}

func main(){
	list := &LinkedList{Value: 0}
	list.Next = &LinkedList{Value: 1}
	list.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	//list.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 10}

	val := NodeSwap(list)
	fmt.Println("val -> ",val)
}

func displayNodes (list *LinkedList) {
	for list != nil {
		fmt.Printf(" %d -> ", list.Value)
		list = list.Next
	}
	
	fmt.Println()
}