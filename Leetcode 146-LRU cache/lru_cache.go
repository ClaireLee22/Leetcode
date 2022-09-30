package main

import (
	"fmt"
	"strings"
)

type LRUCache struct {
	capacity   int
	cache      map[int]*LinkedListNode
	linkedList *LinkedList
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*LinkedListNode)
	return LRUCache{
		capacity:   capacity,
		cache:      cache,
		linkedList: &LinkedList{},
	}
}

func (this *LRUCache) Get(key int) int {
	if _, found := this.cache[key]; !found {
		return -1
	}

	value := this.cache[key].value
	this.linkedList.removeNode(this.cache[key])
	this.linkedList.insertAtTail(key, value)
	this.cache[key] = this.linkedList.getTail()

	return value
}

func (this *LRUCache) Put(key int, value int) {
	if _, found := this.cache[key]; !found {
		if this.linkedList.size >= this.capacity {
			this.linkedList.insertAtTail(key, value)
			this.cache[key] = this.linkedList.getTail()
			delete(this.cache, this.linkedList.getHead().key)
			this.linkedList.removeHead()
		} else {
			this.linkedList.insertAtTail(key, value)
			this.cache[key] = this.linkedList.getTail()
		}
	} else {
		this.linkedList.removeNode(this.cache[key])
		this.linkedList.insertAtTail(key, value)
		this.cache[key] = this.linkedList.getTail()
	}

}

func (this *LRUCache) printLinkedList() string {
	current := this.linkedList.head
	elements := []string{}
	for current != nil {
		elements = append(elements, fmt.Sprintf("%d", current.value))
		current = current.next
	}
	return strings.Join(elements, " -> ")
}

type LinkedListNode struct {
	key   int
	value int
	next  *LinkedListNode
	prev  *LinkedListNode
}

type LinkedList struct {
	size int
	head *LinkedListNode
	tail *LinkedListNode
}

func (l *LinkedList) removeNode(node *LinkedListNode) {
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == l.head {
		l.head = l.head.next
	}
	if node == l.tail {
		l.tail = l.tail.prev
	}

	l.size -= 1
}

func (l *LinkedList) insertAtTail(key, value int) {
	newNode := &LinkedListNode{key: key, value: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}

	l.size += 1
}

func (l *LinkedList) removeHead() {
	l.removeNode(l.head)
}

func (l *LinkedList) getHead() *LinkedListNode {
	return l.head
}

func (l *LinkedList) getTail() *LinkedListNode {
	return l.tail
}

func main() {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	fmt.Println("Put(1, 1)")
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	lruCache.Put(2, 2)
	fmt.Println("\nPut(2, 2)")
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	fmt.Println("\nGet(1)")
	fmt.Println("value: ", lruCache.Get(1))
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	lruCache.Put(3, 3)
	fmt.Println("\nPut(3, 3)")
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	fmt.Println("\nGet(2)")
	fmt.Println("value: ", lruCache.Get(2))
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	lruCache.Put(4, 4)
	fmt.Println("\nPut(4, 4)")
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	fmt.Println("\nGet(1)")
	fmt.Println("value: ", lruCache.Get(1))
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	fmt.Println("\nGet(3)")
	fmt.Println("value: ", lruCache.Get(3))
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())

	fmt.Println("\nGet(4)")
	fmt.Println("value: ", lruCache.Get(4))
	fmt.Println("cache: ", lruCache.cache)
	fmt.Println("doubly linked list", lruCache.printLinkedList())
}

/*
output:
Put(1, 1)
cache:  map[1:0xc00011a000]
doubly linked list 1

Put(2, 2)
cache:  map[1:0xc00011a000 2:0xc00011a040]
doubly linked list 1 -> 2

Get(1)
value:  1
cache:  map[1:0xc00011a0a0 2:0xc00011a040]
doubly linked list 2 -> 1

Put(3, 3)
cache:  map[1:0xc00011a0a0 3:0xc00011a100]
doubly linked list 1 -> 3

Get(2)
value:  -1
cache:  map[1:0xc00011a0a0 3:0xc00011a100]
doubly linked list 1 -> 3

Put(4, 4)
cache:  map[3:0xc00011a100 4:0xc00011a1a0]
doubly linked list 3 -> 4

Get(1)
value:  -1
cache:  map[3:0xc00011a100 4:0xc00011a1a0]
doubly linked list 3 -> 4

Get(3)
value:  3
cache:  map[3:0xc00011a240 4:0xc00011a1a0]
doubly linked list 4 -> 3

Get(4)
value:  4
cache:  map[3:0xc00011a240 4:0xc00011a2a0]
doubly linked list 3 -> 4
*/
