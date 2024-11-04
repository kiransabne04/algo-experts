package main

import "fmt"

// Problem: Given two singly linked lists, find the node at which they intersect.
// Approach: Use two-pointer technique to sync up the ends of the lists.
// List 1: 1 -> 2 -> 3 -> 4 -> 5 -> null
// List 2: 6 -> 4 -> 5 -> null (intersection starts from node 4)
// Expected Output: Node with value 4

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func findIntersectionNode(ll1 *LinkedList, ll2 *LinkedList) int {
	seenMap := make(map[int]bool)

	for ll1 != nil {
		fmt.Printf(" %d -> ", ll1.Value)
		seenMap[ll1.Value] = true
		ll1 = ll1.Next
	}

	fmt.Println("seenMap -> ", seenMap)

	for ll2 != nil {
		fmt.Println("ll2.Value -> ", ll2.Value)
		if _, found := seenMap[ll2.Value]; found {
			return ll2.Value
		}
		ll2 = ll2.Next
	}
	return -1
}

func main() {
	list := &LinkedList{Value: 1}
	list.Next = &LinkedList{Value: 2}
	list.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next = &LinkedList{Value: 5}

	list2 := &LinkedList{Value: 6}
	list2.Next = &LinkedList{Value: 4}
	list2.Next.Next = &LinkedList{Value: 5}

	val := findIntersectionNode(list, list2)
	fmt.Println("val -> ", val)
}
