package main

import (
	"fmt"
)

var aux[]int

func MergeSort(a []int, p int, r int, cmp func(int, int) bool){
	if p==0 && r==(len(a)-1) {
		aux = make([]int, len(a))
	}

	length := r-p + 1
	q := length/2 + p
	fmt.Println("#MergeSort-Start p=", p, "r=", r, "q=", q, "length:", length," a=", a[p:(r+1)])
	if length > 2{
		MergeSort(a, p, q-1, cmp)
		MergeSort(a, q, r, cmp)
	}
	i := p
	j := q
	k := p
	for i<=q || j<r {
		hasLeft := i<q
		hasRight := j<=r
		fmt.Println("#Merge i=", i, "j=", j, "k=", k)
		if hasLeft && hasRight {
			if cmp(a[i], a[j]) {
				aux[k] = a[i]
				i++
			}else{
				aux[k] = a[j]
				j++
			}
		}else if hasLeft {
			aux[k] = a[i]
			i++
		}else if hasRight {
			aux[k] = a[j]
			j++
		}else{
			break
		}
		k++
	}

	for k:=p; k<=r; k++ {
		a[k] = aux[k]
	}
	fmt.Println("#MergeSort-End p=", p, "r=", r, " a=", a[p:(r+1)])
}

func main() {
	a := []int{6, 5, 3, 1, 8, 7, 2, 4}
	MergeSort(a, 0, len(a)-1, func(a int, b int) bool {return a < b})
	fmt.Println("Sorted:", a)
}
