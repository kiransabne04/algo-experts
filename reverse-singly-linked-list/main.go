package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	// Write your code here.
	// 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> nil
	var previousNode *LinkedList
	currentNode := head
	var nextNode *LinkedList
	for currentNode != nil {
		nextNode = currentNode.Next

		currentNode.Next = previousNode

		previousNode = currentNode
		currentNode = nextNode

	}
	return previousNode
}

func main(){
	head := &LinkedList{Value: 0}
	head.Next = &LinkedList{Value: 1}
	head.Next.Next = &LinkedList{Value: 2}
	head.Next.Next.Next = &LinkedList{Value: 3}
	head.Next.Next.Next.Next = &LinkedList{Value: 4}
	head.Next.Next.Next.Next.Next = &LinkedList{Value: 5}

	val := ReverseLinkedList(head)
	fmt.Println(val)
}