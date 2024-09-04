package heap

import (
	// "container/heap"
	"fmt"
)

func maxHeap(a []int, heapSize, i int) {
	l := 2*i
	r := 2*i+1

	largest := i
	if l < heapSize && a[l] > a[i] {
		largest = l
	}

	if r < heapSize && a[r] > a[i] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
	}
	maxHeap(a, heapSize, largest)
}

func buildMaxHeap(a []int) {
	arr := len(a)

	for i := arr/2 - 1; i >= 0; i-- {
		maxHeap(a, arr, i)
	}
}

func main() {
	arr := []int{100, 90, 85, 80, 75, 70, 65, 60, 55, 50, 45, 40}

}
