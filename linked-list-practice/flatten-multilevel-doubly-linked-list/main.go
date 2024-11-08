package main

import "fmt"

// Problem: Given a doubly linked list where each node might have a child list, flatten the list so that all nodes appear in a single level.
// Follow-up: Reverse the flattened multilevel list.
// Input: 1 -> 2 -> 3 -> 4 -> 5 -> null with 3 -> 7 -> 8 -> null as a child of 3
// Expected Output: 1 -> 2 -> 3 -> 7 -> 8 -> 4 -> 5 -> null

// 1 <-> 2 <-> 3 -> child( 7  -> 8 ) <-> 4 <-> 5

type DoublyLinkedList struct {
	Value int
	Next  *DoublyLinkedList
	Prev  *DoublyLinkedList
	Child *DoublyLinkedList
}

func flattenList(head *DoublyLinkedList) *DoublyLinkedList {

	current := head

	for current != nil {
		if current.Child != nil {
			// check for child & change the first part pointers

			//store the nextNode of the current, to append at the end of the child list later to flatten
			nextNode := current.Next

			//childhead
			childHead := flattenList(current.Child)

			//currentNode and childhead node prointers arrangement
			current.Next = childHead
			childHead.Prev = current
			current.Child = nil

			// find tail in child list
			tailPtr := childHead
			for tailPtr.Next != nil {
				tailPtr = tailPtr.Next
			}
			// Connect the tail of the flattened child list to the next node
			if nextNode != nil {
				tailPtr.Next = nextNode
				nextNode.Prev = tailPtr
			}

		}
		current = current.Next
	}
	printList(head)
	printList(reverseDll(head))
	return head

}

func reverseDll(head *DoublyLinkedList) *DoublyLinkedList {

	var temp *DoublyLinkedList
	current := head

	// Traverse the list and swap `Next` and `Prev` for each node
	for current != nil {
		// Swap the `Next` and `Prev` pointers
		temp = current.Prev
		current.Prev = current.Next
		current.Next = temp

		// Move `current` to the next node in the original list, which is `current.Prev` after the swap
		current = current.Prev
	}

	// After the loop, `temp` will be at the last node, which is the new head
	if temp != nil {
		head = temp.Prev
	}

	return head
}

func main() {
	// Main list nodes
	node1 := &DoublyLinkedList{Value: 1}
	node2 := &DoublyLinkedList{Value: 2}
	node3 := &DoublyLinkedList{Value: 3}
	node4 := &DoublyLinkedList{Value: 4}
	node5 := &DoublyLinkedList{Value: 5}

	// Linking main list nodes
	node1.Next = node2
	node2.Prev = node1
	node2.Next = node3
	node3.Prev = node2
	node3.Next = node4
	node4.Prev = node3
	node4.Next = node5
	node5.Prev = node4

	// Step 3: Create the Child List
	// Child list nodes
	node7 := &DoublyLinkedList{Value: 7}
	node8 := &DoublyLinkedList{Value: 8}

	// Linking child list nodes
	node3.Child = node7
	node7.Next = node8
	node8.Prev = node7

	// Display the list to verify
	printList(node1)
	flattenList(node1)
}

func printList(head *DoublyLinkedList) {
	for head != nil {
		fmt.Printf("%d", head.Value)
		if head.Child != nil {
			fmt.Printf(" -> child( %d ", head.Child.Value)
			child := head.Child.Next
			for child != nil {
				fmt.Printf(" -> %d ", child.Value)
				child = child.Next
			}
			fmt.Printf(")")
		}
		head = head.Next
		if head != nil {
			fmt.Printf(" <-> ")
		}
	}
	fmt.Println()
}
