# Contents

* [0. Introduction](#intro)
* [1. Heapsort](#heapsort)
* [2. Quicksort](#quicksort)

# 0. Introduction <i id="intro"></i>

In practice, the numbers to be sorted are rearely isolated values. Each is usually part of a collection of data called a **record**. Each record contains a **key**, which is the value to be sorted. The remainder of the record consists of **satellite data**, which are usually carried around with the key. If each record includes a large amount of satellite data, we often permute an array of pointers to the records rather than the records themselves in order to minimize data movement.

Θ-notation:

Θ(g(n)) = { f(n): there exist positive constants c<sub>1</sub>, c<sub>2</sub>, and n<sub>0</sub> such that 0≤c<sub>1</sub>g(n)≤f(n)≤c<sub>2</sub>g(n) for all n≥n<sub>0</sub> }.

Θ(g(n))表示对于f(n)的上限、下限都可以用g(n)的常数倍表示（常数大于0）

O-notation:

O(g(n)) = { f(n): there exist positive contants c and n<sub>0</sub> such that 0≤f(n)≤cg(n) for all n≥n<sub>0</sub> }.

O(g(n))表示，f(n)的上限一定不会超过g(n)的某个常数倍



Algorithm      | Worst-case running time | Average-case/expected running time
---------------|-------------------------|-----------------------------------
Insertion sort |    Θ(n<sup>2</sup>)     |  Θ(n<sup>2</sup>)
Merge Sort     |    Θ(n lg n)            |  Θ(n lg n)
Heapsort       |    O(n lg n)            |  -
Quicksort      |    Θ(n<sup>2</sup>)     |  Θ(n lg n) (expected)
Counting sort  |    Θ(k + n)             |  Θ(k + n)
Radix sort     |    Θ(d(n + k))          |  Θ(d(n + k))
Bucket sort    |    Θ(n<sup>2</sup>      |  Θ(n)   (average case)


# 1. Heapsort <i id="heapsort"></i>

The **(binary) heap** data structure is an array object ithat we can view as a nearly complete binary tree.

## Getting node from a heap

PARENT(i) = a[i/2]

LEFT(i) = a[2i]

RIGHT(i) = a[2i+1]

**max-heap**: A heap that parent node always larger than its childs.   
**min-heap**: A heap that parent node is smaller than all of its childs.

Height of a heap with n nodes: Θ(lg n)

* The MAX_HEAPIFY procedure, which runs in O(lg n) time, is the key to maintaining the max-heap property.
* The BUILD-MAX-HEAP procedure, which runs in linear time, produces a max-heap from an unordered input array.
* The HEAPSORT procedure, which runs in O(n lg n) time, sorts an array in place.
* The MAX-HEAP-INSERT, HEAP-EXTRACT-MAX, HEAP-INCREASE-KEY, and HEAP-MAXIMUM procedures, which run in O(lg n) time, allow the heap data structure to implement a priority queue.

## Summary

Heapsort is done with three procedures:

1. **Max-Heapify**
   Max-Heapify can ensure that the smallest binary tree unit can be adjusted and the largest element becomes the root

2. **Build-Max-Heap**
   Do max-heapify in bottom-up direction, so when one build-max-heap is done, the largest element will be the root of the tree(First element)

3. **Heap-Sort**
   After one build-max-heap, largest element is at a[0]. Swap a[0] with last element of current array, so tree index will not be broken and easier to calculate the right index.
   Do this until all elements are sorted.

## PriorityQueue

PriorityQueue is built based on a max-heap or a min-heap. Each time we pick the maximum/minimum element from a priority queue, and remaining of them will keep max-heap/min-heap by maxHeapify/minHeapify(O(lgn))

As we just pick the root element each time, a priority queue must be a max-heap/min-heap, but not necessarily be a sorted array just like heap sort.

### Running time of main proedures
* **Extract Max** O(lg n)
* **Get Max** Θ(1)
* **Increase Key** O(lg n)
* **Insert** O(lg n)

# 2. Quicksort <i id="quicksort"></i>

The quicksort algorithm has a worst-case running time of Θ(n<sup>2</sup>) on an input array
of n numbers. Despite this slow worst-case running time, quicksort is often the best
practical choice for sorting because it is remarkably efficient on the average: its
expected running time is Θ(n lg n), and the constant factors hidden in the ‚.n lg n/
notation are quite small. It also has the advantage of sorting in place,
and it works well even in virtual-memory environments.

Q: When does quick result to worst running time?

## Partition

Partition is the key operation in quicksort. Select a pivot element, divide the array into two parts: left elements are smaller than pivot, and right elements are larger.

Psudo code:

    def partition(A, p, r):
        pivot = A[r]
        i = p-1
        for j in range(p, r):
            if A[j] < pivot:
                i += 1
                A[i], A[j] = A[j], A[i]
        A[i+1], A[r] = A[r], A[i+1]
        return i+1

### Randomized partition

Quicksort result to worst case when partition is completely unbalanced - each time 0 element on one side. So a trick is to make it randomized:

    def randomizedPartition(A, p, r):
        i = rand(p, r)
        A[i], A[r] = A[r], A[i]
        return partition(A, p, r)

## Quicksort

    def quicksort(A, p, r):
        if p < r:
            q = partition(A, p, r)
            quicksort(A, p, q-1)
            quicksort(A, q+1, r)
