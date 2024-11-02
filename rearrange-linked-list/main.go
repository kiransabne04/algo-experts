package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RearrangeLinkedList(head *LinkedList, k int) *LinkedList {
	// Write your code here.
	// 3 dummy nodes for 3 part of data
	var smallHead, smallTail *LinkedList
	var equalHead, equalTail *LinkedList
	var greaterHead, greaterTail *LinkedList 
	
	//lessTail, equalTail, greaterTail := lessHead, equalHead, greaterHead //tails to track ending node of the partition.

	current := head
	for current != nil {
		fmt.Println("currentNode -> ", current.Value)
		if current.Value < k {
			
			if smallHead  == nil {
				smallHead = current // Initialize smallHead if it’s the first smaller node
				smallTail = current // smallTail should also point to the first node initially
			} else if smallTail != nil {
				smallTail.Next = current // Append to the end of the smaller list
				smallTail = current // Update smallTail to be the last node in the smaller list
			}
		} else if current.Value == k {
			if equalHead == nil {
				equalHead = current
				equalTail = current
			} else if equalTail != nil {
				equalTail.Next = current
				equalTail = current
			}
		} else if current.Value > k {
			if greaterHead == nil {
				greaterHead = current
				greaterTail = current
			} else if greaterTail != nil {
				greaterTail.Next = current
				greaterTail = current
			}
		}

		previousProcessedNode := current
		current = current.Next
		previousProcessedNode.Next = nil
	}

	displayNodes(smallHead)
	fmt.Println("lessTail -> ", smallTail)
	displayNodes(equalHead)
	fmt.Println("equalTail -> ", equalTail)
	displayNodes(greaterHead)
	fmt.Println("greaterTail -> ", greaterTail)

	// combine the list & determine new head
	if smallTail != nil {
		smallTail.Next = equalHead // Connect less list to equal list
	} else {
		smallHead = equalHead // If no less list, start with equal list
	}

	if equalTail != nil {
		equalTail.Next = greaterHead // Connect equal list to greater list
	} else if smallTail != nil {
		smallTail.Next = greaterHead // If no equal list, connect less directly to greater
	}

	// Determine the new head
	if smallHead != nil {
		return smallHead // Start with less list if it exists
	} else if equalHead != nil {
		return equalHead // Start with equal list if no less list
	} else {
		return greaterHead // Otherwise, start with greater list
	}

}

func displayNodes(list *LinkedList) {
	for list != nil {
		fmt.Printf(" %d -> ", list.Value)
		list = list.Next
	}
	fmt.Println()
}

func main () {

	list := &LinkedList{Value: 3}
	list.Next = &LinkedList{Value: 0}
	list.Next.Next = &LinkedList{Value: 5}
	list.Next.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next.Next = &LinkedList{Value: 1}
	list.Next.Next.Next.Next.Next = &LinkedList{Value: 4}
	
	val := RearrangeLinkedList(list, 3)

	fmt.Println("val -> ", val)
}