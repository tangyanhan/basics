package main

import "math/rand"
import (
	"time"
	"fmt"
	"strings"
)

func randomizeData(data []int){
	rand.Seed(time.Now().Unix())

	L := len(data)
	for idx:=0; idx < L; idx++ {
		swapIdx := rand.Intn(L)
		data[idx], data[swapIdx] = data[swapIdx], data[idx]
	}
}

func parentIndex(i int) int {
	if i%2 == 0 { // Right node or parent node
		return (i - 2)/2
	}
	// Left node
	return (i - 1)/2
}

// Array start from 0, so it's a bit different.
func leftChildIndex(i int) int {
	return 2 * i + 1
}

func rightChildIndex(i int) int {
	return 2 * i + 2
}

func printHeap(heap []int, i int, level int) {
	if i == 0 {
		fmt.Println("=====Heap")
	}

	if i>= len(heap) { return }
	// Print parent
	prefix := strings.Repeat(" ", 4 * level) + strings.Repeat("-", 4)
	fmt.Println(prefix, heap[i])

	// Print children if any
	left := leftChildIndex(i)
	printHeap(heap, left, level+1)

	right := rightChildIndex(i)
	printHeap(heap, right, level+1)

	if i == 0 {
		fmt.Println("=====Heap End")
	}
}

func maxHeapify(heap []int, root int){
	largestIdx := root
	left := leftChildIndex(root)
	if left<len(heap) && heap[left] > heap[largestIdx] {
		largestIdx = left
	}
	right := rightChildIndex(root)
	if right<len(heap) && heap[right] > heap[largestIdx] {
		largestIdx = right
	}
	if largestIdx != root {
		heap[root], heap[largestIdx] = heap[largestIdx], heap[root]
	}
}

func getMaximum(queue []int) int {
	if len(queue) == 0 {
		panic("Emtpy queue")
	}
	return queue[0]
}

func extractMax(queue[]int) (int, []int){
	if len(queue) == 0 {
		panic("Queue underflow")
	}
	max := queue[0]
	queue[0] = queue[len(queue) - 1]
	queue = queue[:(len(queue) - 1)]  // Shrink length of queue
	if len(queue) > 1 {
		maxHeapify(queue, 0)
	}
	return max, queue
}

func increaseKey(queue []int, i int, key int) {
	if i >= len(queue) {
		panic("Index too large for the queue")
	}
	if key < queue[i] {
		panic("Value to increase is smaller than original")
	}
	queue[i] = key
	for parent:= parentIndex(i); i>0 && queue[parent]<queue[i]; i=parent {
		queue[parent], queue[i] = queue[i], queue[parent]
	}
}

func insert(queue[]int, key int) []int{
	queue = append(queue, -99999999)
	increaseKey(queue, len(queue)-1, key)
	return queue
}

func main() {
	data := [5]int{5, 7, 9, 1, 2}
	queue := make([]int, 0, 5)

	for i, v := range data {
		queue = insert(queue, v)
		fmt.Println("#", i, " Inserted value=", v, "queue", queue)
		printHeap(queue, 0, 0)
	}

	var max int
	for i:=0; i<5; i++ {
		max, queue = extractMax(queue)
		fmt.Println("Extracted:", max, " queue", queue)
	}
}