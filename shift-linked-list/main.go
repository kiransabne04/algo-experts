package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	// Write your code here.

	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// find the length & tail of the list
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}
	fmt.Println("tail -> ", tail)

	//adjust k based on its sign
	if k > 0 {
		k = k % length
	} else if k < 0 {
		// k = (length + k) % length
		k = (k % length + length) % length
	}

	fmt.Println("k -> ", k)

	if k == 0 {
		return head
	}

	
	// find the new tail & new head
	newTailPosition := length - k - 1
	newTail := head
	fmt.Println("newTailPosition ", newTailPosition, "newTail ", newTail)

	for i := 0; i < newTailPosition; i++{
		newTail = newTail.Next
	}
	newHead := newTail.Next
	fmt.Println("newHead -< ", newHead)

	// break & reatach list
	newTail.Next = nil
	tail.Next = head

	fmt.Println("newTail.Next && tail.Next -< ", newTail.Next, tail.Next)

	return newHead
}

func main(){
	list := &LinkedList{Value: 0}
	list.Next = &LinkedList{Value: 1}
	list.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next.Next = &LinkedList{Value: 5}

	val := ShiftLinkedList(list, -8)
	fmt.Println("val -> ", val)
}