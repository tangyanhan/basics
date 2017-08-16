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

func buildMaxHeap(heap []int) {
    bottomRootIdx := parentIndex(len(heap)-1)
    fmt.Println("Build max heap from", bottomRootIdx, " Heap size=", len(heap))
    for i:=bottomRootIdx; i>=0; i-- {
        maxHeapify(heap, i)
    }
}

func heapSort(heap[] int) {
    buildMaxHeap(heap)
    for i:=len(heap)-1; i>=1; i-- {
        heap[i], heap[0] = heap[0], heap[i]
        buildMaxHeap(heap[:i])
        fmt.Println("#",i, " BuildMaxHeap")
        printHeap(heap[:i], 0, 0)
    }
}

func main() {
    var x = [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    data := x[:]
    randomizeData(data)
    fmt.Println("Randomized array:", data)

    fmt.Println("Original heap=============")
    printHeap(data, 0, 0)
    heapSort(data)
    fmt.Println("Sorted heap===============")
    printHeap(data, 0, 0)
    fmt.Println("Sorted array:", data)
}