package edsort

import "sort"

// InsertionSort picks an element from the slice, and inserts it in
// its proper position.
func InsertionSort(data sort.Interface) {
	for i := 1; i < data.Len(); i++ {
		j := i
		for j > 0 && data.Less(j, j-1) {
			data.Swap(j, j-1)
			j--
		}
	}
}
