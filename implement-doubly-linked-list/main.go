package main

import "fmt"

// task is to implement doubly linked list
type Node struct {
	ID         string
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{Head: nil, Tail: nil}
}

func (ll *DoublyLinkedList) setHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.insertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) setTail(node *Node) {
	if ll.Tail == nil {
		ll.setHead(node)

		return
	}
	ll.insertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) printForward() {
	node := ll.Head
	for node != nil {
		fmt.Println("%d <-> ", node.Value)
		node = node.Next
	}
	fmt.Println(nil)
}

func (ll *DoublyLinkedList) printBackward() {
	node := ll.Tail
	for node != nil {
		fmt.Println("%d <-> ", node.Value)
		node = node.Prev
	}
}

// InsertBefore inserts the nodeToInsert before the given node
func (ll *DoublyLinkedList) insertBefore(node, nodeToInsert *Node) {
	// 10 <> 20 <> 30 <> 40 <> 50 <> 60 <> 70 | i want to add node 100 before node 30
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	ll.remove(node)

	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}

	node.Prev = nodeToInsert

}

// insertAtPosition inserts nodeToInsert at the specified position.
func (ll *DoublyLinkedList) insertAtPosition(position int, inputNode *Node) {
	// 10 <> 20 <> 30 <> 40 <> 50 <> 60 <> 70 | i want to add node 100 at 3rd position

	if position == 1 {
		ll.setHead(inputNode)
		return
	}

	node := ll.Head
	currentPosition := 1

	for node != nil && currentPosition != position {
		node = node.Next
		currentPosition++
	}
	fmt.Println("node & currentPosition ", node.Value, currentPosition, position)
	// now at the posiution
	if node != nil {
		ll.insertBefore(node, inputNode)
	} else {
		ll.setTail(inputNode)
	}
}

// insertAfter inserts the new node after given node
func (ll *DoublyLinkedList) insertAfter(node, inputNode *Node) {
	// 10 <> 20 <> 30 <> 40 <> 50 <> 60 <> 70 | i want to add node 100 after node 30
	if inputNode == ll.Head && inputNode == ll.Tail {
		return
	}
	ll.remove(node)

	inputNode.Prev = node
	inputNode.Next = node.Next

	if node.Next == nil {
		ll.Tail = node
	} else {
		node.Next.Prev = inputNode
	}
	node.Next = inputNode
}

// removeNodesWithValue removes all nodes with the specified value.
func (dll *DoublyLinkedList) removeNodesWithValue(value int) {
	node := dll.Head
	for node != nil {
		nextNode := node.Next
		if node.Value == value {
			dll.remove(node)
		}
		node = nextNode
	}
}

//Remove removes the specified node from linkedlist
func (ll *DoublyLinkedList) remove(node *Node) {
	// is node is head
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	// if node is tail
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	ll.unlinkNode(node)
}

// containsNodeWithValue checks if a node with the given value exists in the list.
func (ll *DoublyLinkedList) containsNodeWithValue(value int) bool {
	node := ll.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

//unlink Node disconnects the node from the linkedlist
func (ll *DoublyLinkedList) unlinkNode(node *Node) {
	// if node.PreviousNode is not null then point the node's prev next pointer to node's next
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Next = nil
	node.Prev = nil
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
