package main

import "fmt"

// Detect and Remove Cycle in a Linked List
// Input: 1 -> 2 -> 3 -> 4 -> 5 -> 3 (cycle starts again at node 3)
// Expected Output: true (for detection) and 1 -> 2 -> 3 -> 4 -> 5 -> null (for removal)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func findCycle(head *LinkedList) *LinkedList {

	slowPtr := head
	fastPtr := head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if slowPtr == fastPtr {
			break
		}
	}

	fmt.Println("slowPtr & fastPtr ", slowPtr, fastPtr)

	if fastPtr == nil || fastPtr.Next == nil {
		// no cycle detected
		return nil
	}
	fmt.Println("slowPtr & fastPtr ", slowPtr, fastPtr)
	slowPtr = head
	fmt.Println("slowPtr ", slowPtr, fastPtr)
	for slowPtr != fastPtr {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}
	fmt.Println("slowPtr & fastPtr ", slowPtr, fastPtr)

	cycleStart := slowPtr
	fmt.Println("cyclestart -< ", cycleStart)
	ptr := cycleStart
	fmt.Println("ptr -< ", ptr)
	for ptr.Next != cycleStart {

		ptr = ptr.Next
		fmt.Println("ptr -<, cycleStart -> ", ptr, cycleStart)
	}
	fmt.Println("slowPtr & fastPtr ", slowPtr, fastPtr)
	ptr.Next = nil
	fmt.Println("slowPtr & fastPtr ", slowPtr, fastPtr)
	//displayNodes(head)

	return head
}

func displayNodes(list *LinkedList) {
	for list != nil {
		fmt.Printf(" %d-> ", list.Value)
		list = list.Next
	}
	fmt.Println()
}

func main() {

	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 2}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next = &LinkedList{Value: 5}
	list.Next.Next.Next.Next.Next = list.Next.Next
	//list.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 4}
	//list.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 5}

	val := findCycle(list)

	fmt.Println("val -> ", val)
}
