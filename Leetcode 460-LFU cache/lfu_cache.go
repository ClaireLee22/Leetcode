package main

import (
	"fmt"
	"strings"
)

type LFUCache struct {
	capacity     int
	size         int
	minFreq      int
	cache        map[int]*LinkedListNode
	frequencyMap map[int]*LinkedList
}

func Constructor(capacity int) LFUCache {
	cache := make(map[int]*LinkedListNode)
	frequencyMap := make(map[int]*LinkedList)
	return LFUCache{
		capacity:     capacity,
		size:         0,
		minFreq:      0,
		cache:        cache,
		frequencyMap: frequencyMap,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, found := this.cache[key]; !found {
		return -1
	}

	tempNode := this.cache[key]
	this.frequencyMap[tempNode.freq].removeNode(tempNode)
	if this.frequencyMap[tempNode.freq].head == nil {
		delete(this.frequencyMap, tempNode.freq)
		if this.minFreq == tempNode.freq {
			this.minFreq += 1
		}
	}

	this.cache[key].freq += 1
	this.initLinkedList(this.cache[key].freq)
	this.frequencyMap[this.cache[key].freq].insertAtTail(this.cache[key])
	return this.cache[key].value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if _, found := this.cache[key]; found {
		this.Get(key)
		this.cache[key].value = value
		return
	}

	if this.size == this.capacity {
		delete(this.cache, this.frequencyMap[this.minFreq].head.key)
		this.frequencyMap[this.minFreq].removeHead()
		if this.frequencyMap[this.minFreq].head == nil {
			delete(this.frequencyMap, this.minFreq)
		}
		this.size -= 1
	}

	this.minFreq = 1
	this.cache[key] = &LinkedListNode{
		key:   key,
		value: value,
		freq:  this.minFreq,
	}

	this.initLinkedList(this.cache[key].freq)
	this.frequencyMap[this.cache[key].freq].insertAtTail(this.cache[key])
	this.size += 1
}

func (this *LFUCache) initLinkedList(freq int) {
	if _, found := this.frequencyMap[freq]; !found {
		this.frequencyMap[freq] = &LinkedList{}
	}
}

func (this *LFUCache) printFrequencyMap() {
	fmt.Println("frequencyMap: ")
	for freq, linkedList := range this.frequencyMap {
		fmt.Println("frequency: ", freq)
		fmt.Println("linked list", printLinkedList(linkedList))
	}
}

func printLinkedList(linkedList *LinkedList) string {
	current := linkedList.head
	nodes := []string{}
	for current != nil {
		nodes = append(nodes, fmt.Sprintf("%d", current.value))
		current = current.next
	}
	return strings.Join(nodes, " -> ")
}

type LinkedListNode struct {
	key   int
	value int
	freq  int
	next  *LinkedListNode
	prev  *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func (ll *LinkedList) insertAtTail(node *LinkedListNode) {
	node.prev, node.next = nil, nil
	if ll.tail == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		node.prev = ll.tail
		ll.tail = node
	}
}

func (ll *LinkedList) removeNode(node *LinkedListNode) {
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == ll.head {
		ll.head = ll.head.next
	}
	if node == ll.tail {
		ll.tail = ll.tail.prev
	}
}

func (ll *LinkedList) removeHead() {
	ll.removeNode(ll.head)
}

func main() {
	lfuCache := Constructor(2)
	lfuCache.Put(1, 1)
	fmt.Println("Put(1, 1)")
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	lfuCache.Put(2, 2)
	fmt.Println("\nPut(2, 2)")
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	fmt.Println("\nGet(1)")
	fmt.Println("value", lfuCache.Get(1))
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	lfuCache.Put(3, 3)
	fmt.Println("\nPut(3, 3)")
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	fmt.Println("\nGet(2)")
	fmt.Println("value", lfuCache.Get(2))
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	fmt.Println("\nGet(3)")
	fmt.Println("value", lfuCache.Get(3))
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	lfuCache.Put(4, 4)
	fmt.Println("\nPut(4, 4)")
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	fmt.Println("\nGet(1)")
	fmt.Println("value", lfuCache.Get(1))
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	fmt.Println("\nGet(3)")
	fmt.Println("value", lfuCache.Get(3))
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()

	fmt.Println("\nGet(4)")
	fmt.Println("value", lfuCache.Get(4))
	fmt.Println("cache: ", lfuCache.cache)
	lfuCache.printFrequencyMap()
}

/* output:
Put(1, 1)
cache:  map[1:0xc0000701e0]
frequencyMap:
frequency:  1
linked list 1

Put(2, 2)
cache:  map[1:0xc0000701e0 2:0xc000070240]
frequencyMap:
frequency:  1
linked list 1 -> 2

Get(1)
value 1
cache:  map[1:0xc0000701e0 2:0xc000070240]
frequencyMap:
frequency:  1
linked list 2
frequency:  2
linked list 1

Put(3, 3)
cache:  map[1:0xc0000701e0 3:0xc000070390]
frequencyMap:
frequency:  1
linked list 3
frequency:  2
linked list 1

Get(2)
value -1
cache:  map[1:0xc0000701e0 3:0xc000070390]
frequencyMap:
frequency:  2
linked list 1
frequency:  1
linked list 3

Get(3)
value 3
cache:  map[1:0xc0000701e0 3:0xc000070390]
frequencyMap:
frequency:  2
linked list 1 -> 3

Put(4, 4)
cache:  map[3:0xc000070390 4:0xc000070570]
frequencyMap:
frequency:  1
linked list 4
frequency:  2
linked list 3

Get(1)
value -1
cache:  map[3:0xc000070390 4:0xc000070570]
frequencyMap:
frequency:  1
linked list 4
frequency:  2
linked list 3

Get(3)
value 3
cache:  map[3:0xc000070390 4:0xc000070570]
frequencyMap:
frequency:  1
linked list 4
frequency:  3
linked list 3

Get(4)
value 4
cache:  map[3:0xc000070390 4:0xc000070570]
frequencyMap:
frequency:  2
linked list 4
frequency:  3
linked list 3
*/
