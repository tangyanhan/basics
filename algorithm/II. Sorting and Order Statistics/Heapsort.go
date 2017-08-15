package main

import "fmt"
import "strings"
import "math/rand"
import "time"

func randomizeData(data []int){
    rand.Seed(time.Now().Unix())

    L := len(data)
    for idx:=0; idx < L; idx++ {
        swapIdx := rand.Intn(L)
        data[idx], data[swapIdx] = data[swapIdx], data[idx]
    }
}

// Util functions
func parentIndex(i int) int {
    return i/2 - 1
}

// Array start from 0, so it's a bit different.
func leftChildIndex(i int) int {
    return 2 * i + 1
}

func rightChildIndex(i int) int {
    return 2 * i + 2
}

func printHeap(heap []int, i int, level int) {
    if i>= len(heap) { return }
    // Print parent
    prefix := strings.Repeat(" ", 4 * level) + strings.Repeat("-", 4)
    fmt.Println(prefix, heap[i])

    // Print children if any
    left := leftChildIndex(i)
    if left < len(heap) {
        printHeap(heap, left, level+1)
    }
    right := rightChildIndex(i)
    if right < len(heap) {
        printHeap(heap, right, level+1)
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

func buildMaxHeap(heap []int) {
    bottomRootIdx := parentIndex(len(heap)-1)
    for i:=bottomRootIdx; i>=0; i-- {
        maxHeapify(heap, i)
    }
}

func heapSort(heap[] int) {
    buildMaxHeap(heap)
    printHeap(heap, 0, 0)
    for i:=len(heap)-1; i>=1; i-- {
        heap[i], heap[0] = heap[0], heap[i]
        buildMaxHeap(heap[:i])
    }
}

func main() {
    var x = [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    data := x[:]
    randomizeData(data)

    fmt.Println("Original heap=============")
    printHeap(data, 0, 0)
    heapSort(data)
    fmt.Println("Sorted heap===============")
    printHeap(data, 0, 0)
    fmt.Println("Sorted array:", data)
}