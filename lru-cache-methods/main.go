package main

import "fmt"

type Node struct {
	Key string
	Value int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	maxSize int
	cache map[string]*Node
	Head *Node
	Tail *Node
}

// <HEAD> <1> <2> |<3>| <4> <5> <6> <TAIL>

// func NewLRUCache(maxSize int) *LRUCache {
// 	lru := &LRUCache {
// 		maxSize: maxSize,
// 		cache: map[string]*Node{},
// 	}
	
// 	lru.Head = &Node{}
// 	lru.Tail = &Node{}

// 	return lru
// }

// helper functions for add a node, remove a node, move node to head to mark it as recently used
func (lru *LRUCache) remove (node *Node) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

// new node is added then it will also be most recently used node
func (lru *LRUCache) add (nodeToInsert *Node) {
// <HEAD> <1> <2> <3> <4> <5> <6> <TAIL>
	
	nodeToInsert.Prev = lru.Head
	nodeToInsert.Next = lru.Head.Next

	lru.Head.Next.Prev = nodeToInsert
	lru.Head.Next = nodeToInsert

}

func (lru *LRUCache) moveToHead (node *Node) {
	lru.remove(node)
	lru.add(node)
}

func NewLRUCache(size int) *LRUCache {
	// Write your code here.
	lru := &LRUCache{
		maxSize: size,
		cache: make(map[string]*Node),
	}
	lru.Head = &Node{}
	lru.Tail = &Node{}

	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head

	return lru
}

func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	// Write your code here.
	if nodeVal, found := cache.cache[key]; found {
		nodeVal.Value = value
		cache.moveToHead(nodeVal)
	} else {
		newNode := &Node{Key: key, Value: value}
		cache.cache[key] = newNode
		cache.add(newNode)

		if cache.maxSize < len(cache.cache) {
			tail := cache.Tail.Prev
			cache.remove(tail)
			delete(cache.cache, tail.Key)

		}
	}
}

// The second return value indicates whether or not the key was found
// in the cache.
func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	// Write your code here.
	if nodeVal, found := cache.cache[key]; found {
		cache.moveToHead(nodeVal)
		return nodeVal.Value, found
	}
	return -1, false
}

// The second return value is false if the cache is empty.
func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	// Write your code here.
	if cache.Head != cache.Tail {
		return cache.Head.Next.Key, true
	}

	return "", false
}


func main() {
	TestLRUCache()
}

func TestLRUCache() {
	cache := NewLRUCache(2)

	// Test case 1: Get key from empty cache
	fmt.Println(cache.Get(1)) // Expected -1 (not found)

	// Test case 2: Add keys and get them
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Expected 1

	// Test case 3: Exceeding capacity (evicting least recently used)
	cache.Put(3, 3)           // Evicts key 2
	fmt.Println(cache.Get(2)) // Expected -1 (key 2 was evicted)

	// Test case 4: Add more keys and check eviction policy
	cache.Put(4, 4)           // Evicts key 1
	fmt.Println(cache.Get(1)) // Expected -1 (key 1 was evicted)
	fmt.Println(cache.Get(3)) // Expected 3 (key 3 is present)
	fmt.Println(cache.Get(4)) // Expected 4 (key 4 is present)

	// Test case 5: Adding a key that already exists
	cache.Put(4, 44)
	fmt.Println(cache.Get(4)) // Expected 44 (value updated)
}
