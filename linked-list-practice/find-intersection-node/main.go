package main

// Problem: Given two singly linked lists, find the node at which they intersect.
// Approach: Use two-pointer technique to sync up the ends of the lists.
// List 1: 1 -> 2 -> 3 -> 4 -> 5 -> null
// List 2: 6 -> 4 -> 5 -> null (intersection starts from node 4)
// Expected Output: Node with value 4