package main

import "fmt"
import "math"
import "container/list"

func BucketSort(a []int, bucketNum int, min int, max int){
	buckets := make([]list.List, bucketNum)
	bucketLength := math.Ceil(float64(max - min)/float64(bucketNum))
	for i:=0; i<len(a); i++ {
		ib := int(math.Floor(float64(a[i] - min)/bucketLength))
		l := &buckets[ib]

		e := (*l).Front()
		for ; e!=nil; e=e.Next() {
			if e.Value.(int) > a[i] {
				(*l).InsertBefore(a[i], e)
				break
			}
		}
		if e==nil {
			(*l).PushBack(a[i])
		}
	}

	for b:=0; b<bucketNum; b++ {
		l := buckets[b]
		fmt.Println("=========Bucket #", b)
		for e:=l.Front(); e!=nil; e=e.Next() {
			fmt.Println(e.Value.(int))
		}
	}
	i := 0
	for b:=0; b<bucketNum; b++ {
		l := buckets[b]
		for e:=l.Front(); e!=nil; e=e.Next() {
			a[i] = e.Value.(int)
			i++
		}
	}
}

func main() {
	a := []int{853, 817, 466, 56, 341, 622, 27, 256, 281, 208, 783, 292, 834, 627, 498, 980, 724, 692, 765, 984, 261, 724, 818, 848, 546, 159, 229, 825, 704, 18}
	BucketSort(a, 10, 0, 1000)
	fmt.Println("Bucket sorted:", a)
}
