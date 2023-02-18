package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || this.maxHeap.Top() >= num {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}

	this.balance()
}

func (this *MedianFinder) balance() {
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
	} else if this.maxHeap.Len() < this.minHeap.Len() {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.Top())
	} else {
		return float64(float64(this.maxHeap.Top()+this.minHeap.Top()) / 2.0)
	}
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h MaxHeap) Top() int { return h[0] }

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h MinHeap) Top() int { return h[0] }

func main() {
	dataSteam := []int{2, 3, 7, 9}
	medianFinder := Constructor()

	for _, num := range dataSteam {
		fmt.Printf("add number: %d\n", num)
		medianFinder.AddNum(num)
		fmt.Println("max heap", medianFinder.maxHeap)
		fmt.Println("min heap", medianFinder.minHeap)
		fmt.Println("median", medianFinder.FindMedian())
	}
}

/*
output:
add number: 2
max heap &[2]
min heap &[]
median 2

add number: 3
max heap &[2]
min heap &[3]
median 2.5

add number: 7
max heap &[3 2]
min heap &[7]
median 3

add number: 9
max heap &[3 2]
min heap &[7 9]
median 5
*/
