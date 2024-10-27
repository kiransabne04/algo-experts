package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	// 
	dummy := &LinkedList{Value: 0}
	current := dummy
	carryForward := 0

	for linkedListOne != nil || linkedListTwo != nil ||  carryForward != 0{
		val1, val2 := 0, 0
		
		if linkedListOne != nil {
			val1 = linkedListOne.Value
			linkedListOne = linkedListOne.Next
		}

		if linkedListTwo != nil {
			val2 = linkedListTwo.Value
			linkedListTwo = linkedListTwo.Next
		}

		sum := val1 + val2 + carryForward
		carryForward = sum / 10
		newVal := sum % 10
		fmt.Println("sum & newVal & carryForward => ", sum, carryForward, newVal)

		current.Next = &LinkedList{Value: newVal}
		current = current.Next
	}

	return dummy.Next
}

func main() {
	// Input linked list 1: 2 -> 4 -> 7 -> 1 (represents 1742 in reverse order)
	l1 := &LinkedList{Value: 2, Next: &LinkedList{Value: 4, Next: &LinkedList{Value: 7, Next: &LinkedList{Value: 1}}}}

	// Input linked list 2: 9 -> 4 -> 5 (represents 549 in reverse order)
	l2 := &LinkedList{Value: 9, Next: &LinkedList{Value: 4, Next: &LinkedList{Value: 5}}}

	val := SumOfLinkedLists(l1, l2)
	fmt.Println(val)
}