package main

import "fmt"

type Node struct {
	ID string
	Value int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}


func NewDoublyLinkedList() *DoublyLinkedList {
	// Write your code here.
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	// Write your code here.
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
	
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	// Write your code here.
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)

}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	// Write your code here.
	// 10 <> 20 <> (100) <> 30 <> 40 <> 50 <> 60 <> 70 | i want to add node 100 before node 30
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return 
	}

	ll.Remove(nodeToInsert)

	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node

	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}

	node.Prev = nodeToInsert

}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	// Write your code here.
	// 10 <> 20 <> 30 <> 40 <> 50 <> 60 <> 70 | i want to add node 100 after node 30
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)

	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next

	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		// if node.Next is not nil
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert

}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	// Write your code here.
	// 10 <> 20 <> 30 <> 40 <> 50 <> 60 <> 70 | i want to add node 100 at 3rd position

	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}

	node := ll.Head
	currentPosition := 1
	
	for node != nil && currentPosition != position {
		node = node.Next
		currentPosition++
	}
	// now at that position
	if node != nil {
		ll.InsertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	// Write your code here.
	node := ll.Head
	for node != nil {
		nextNode := node.Next
		if node.Value == value {
			ll.Remove(node)
		}
		node = nextNode
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	// Write your code here.
	// 10 <> 20 <> 30 <> 40 <> 50 <> 60 <> 70 <> 80 <> 90 // want to delete Node 30
	if node == ll.Head{
		ll.Head = ll.Head.Next
	}

	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}

	ll.unlinkNode(node)

}

func (ll *DoublyLinkedList) unlinkNode(node *Node) {
	// unlink node's prev & next if any. And then set node to  nil
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	node.Next = nil
	node.Prev = nil
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	// Write your code here.
	return false
}

func (ll *DoublyLinkedList) printForward() {
	node := ll.Head

	for node != nil {
		fmt.Sprintf(" <-> %d", node.Value)
		node = node.Next
	}
}

func (ll *DoublyLinkedList) printBackward() {
	node := ll.Head

	for node != nil {
		fmt.Sprintf(" <-> %d", node.Value)
		node = node.Prev
	}
}


// 10 <> 20 <> 30 <> 40 <> 50 <> 60 <> 70

func main() {
	// Create nodes
	nodes := map[string]*Node{
		"1":   {ID: "1", Value: 1},
		"2":   {ID: "2", Value: 2},
		"3":   {ID: "3", Value: 3},
		"3-2": {ID: "3-2", Value: 3},
		"3-3": {ID: "3-3", Value: 3},
		"4":   {ID: "4", Value: 4},
		"5":   {ID: "5", Value: 5},
		"6":   {ID: "6", Value: 6},
	}

	// Initialize the doubly linked list
	dll := NewDoublyLinkedList()

	// Execute the specified operations
	dll.setHead(nodes["5"])
	dll.setHead(nodes["4"])
	dll.setHead(nodes["3"])
	dll.setHead(nodes["2"])
	dll.setHead(nodes["1"])
	dll.setHead(nodes["4"])                   // Move "4" to head again
	dll.setTail(nodes["6"])                   // Set "6" as the tail
	dll.insertBefore(nodes["6"], nodes["3"])  // Insert "3" before "6"
	dll.insertAfter(nodes["6"], nodes["3-2"]) // Insert "3-2" after "6"
	dll.insertAtPosition(1, nodes["3-3"])     // Insert "3-3" at position 1
	dll.removeNodesWithValue(3)               // Remove all nodes with value 3
	dll.remove(nodes["2"])                    // Remove node "2"
	containsFive := dll.containsNodeWithValue(5)

	// Print the linked list
	fmt.Println("Doubly Linked List (Forward):")
	dll.printForward()
	fmt.Println("Doubly Linked List (Backward):")
	dll.printBackward()

	// Check for value existence
	fmt.Println("Contains value 5:", containsFive)
}
