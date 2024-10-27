package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {
	list := &LinkedList{Value: 0}
	list.Next = &LinkedList{Value: 1}
	list.Next.Next = &LinkedList{Value: 2}
	list.Next.Next.Next = &LinkedList{Value: 3}
	list.Next.Next.Next.Next = &LinkedList{Value: 4}
	list.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	list.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	list.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 7}
	list.Next.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 8}
	list.Next.Next.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 9}
	list.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = nil

	list.printForward()

	RemoveKthNodeFromEnd(list, 10)
	//fmt.Println(val)

	list.printForward()
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	// Write your code here.+

	// make 2 pointers currentNode & nextNode, both starting at head of linked list.
	// in first loop we will move only nextNode till the kth position.
	// after that, the currentNode will be still at head & nextNode will be at kth position from start.
	// now we will move both pointer in the new loop. so when nextNode reaches end, the currentNode will be pointing to kth element from end
	// we need thrid variable to track the previousNode of the current. 
	// and later at the currentNode position needed, change the previousNode.Next to currentNode.Next

	currentNode, nextNode := head, head
	var previousNode *LinkedList
	i := 0
	for i < k && nextNode != nil {
		if nextNode == nil {
			return
		}
		nextNode = nextNode.Next
		i++
	}

	// edger case if nextNode is nil, k is the length of the list, so remove the head
	if nextNode == nil {
		*head = *(head).Next
		return
	}

	//fmt.Println("currentNode, nextNode -> ", currentNode.Value, nextNode.Value, previousNode)
	fmt.Println("i & K -> ", i, k )
	for nextNode != nil {
		previousNode = currentNode
		nextNode = nextNode.Next
		currentNode = currentNode.Next
	}
    fmt.Println("currentNode2, previousNode -> ", currentNode.Value, previousNode)

	// set the previousNode pointert to currentNode.Next
	// `previousNode` is now just before the node to remove
	if previousNode != nil {
		previousNode.Next = currentNode.Next
	}

}

func (ll *LinkedList) printForward () {
	for ll != nil {
		fmt.Printf(" %d -> ", ll.Value)
		ll = ll.Next
	}
	fmt.Println()
}