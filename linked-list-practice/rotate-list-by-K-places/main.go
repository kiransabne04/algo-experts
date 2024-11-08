package main

import "fmt"

// Input: 1 -> 2 -> 3 -> 4 -> 5 -> null, k = 2
// Expected Output: 4 -> 5 -> 1 -> 2 -> 3 -> null

type LinkedList struct {
	Value int
	Next *LinkedList
}

func rotateKPlaces(head *LinkedList, k int) *LinkedList {

	current := head
	length := 1

	for current.Next != nil {
		current = current.Next
		length++
	}

	current.Next = head

	if k > 0 {
		k = k % length
	} else if k <  0 {
		k = (k % length + length) % length
	}

	fmt.Println("k -> ", k, current, length)

	newTailPosition := length - k - 1
	newTail := head
	fmt.Println("newTailPosition ", newTailPosition, "newTail ", newTail)

	for i := 0; i < newTailPosition; i++ {
		newTail = newTail.Next
	}

	fmt.Println("newTailPosition ", newTailPosition, "newTail ", newTail)

	newHead := newTail.Next

	fmt.Println("newHead -> ", newHead)
	newTail.Next = nil
	//current.Next = head

	displayNodes(newHead)
	return newHead
}

func displayNodes(head *LinkedList) {
	for head != nil {
		fmt.Printf("%d -> ", head.Value)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 2}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next = &LinkedList{Value: 5}

	val := rotateKPlaces(list, 2)

	fmt.Println(val)
}