package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	// Write your code here.
	// set starting 
	var mergedList *LinkedList
}

func main() {

	headOne := &LinkedList{Value: 2}
	headOne.Next = &LinkedList{Value: 6}
	headOne.Next.Next = &LinkedList{Value: 7}
	headOne.Next.Next.Next = &LinkedList{Value: 8}

	headTwo := &LinkedList{Value: 1}
	headTwo.Next = &LinkedList{Value: 3}
	headTwo.Next.Next = &LinkedList{Value: 4}
	headTwo.Next.Next.Next = &LinkedList{Value: 5}
	headTwo.Next.Next.Next.Next = &LinkedList{Value: 9}
	headTwo.Next.Next.Next.Next.Next = &LinkedList{Value: 10}

	val := MergeLinkedLists(headOne, headTwo)
	fmt.Println("val -> ", val)

}
