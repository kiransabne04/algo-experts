package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {
	// Create a linked list with a loop for testing: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 3
	head := &LinkedList{Value: 0}
	head.Next = &LinkedList{Value: 1}
	head.Next.Next = &LinkedList{Value: 2}
	head.Next.Next.Next = &LinkedList{Value: 3}
	head.Next.Next.Next.Next = &LinkedList{Value: 4}
	head.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next // Creates a cycle (6 -> 3)

	// Check if the linked list has a cycle
	_ = FindLoop(head)
}
func FindLoop(head *LinkedList) *LinkedList {
	slow, fast := head, head

	// Step 1: Detect if there is a cycle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// Cycle detected
			break
		}
	}

	// If no cycle is detected, return nil
	if fast == nil || fast.Next == nil {
		return nil
	}

	// Step 2: Find the start of the cycle
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	// The point where slow and fast meet is the start of the loop
	return slow
}