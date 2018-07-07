package sort

import (
	"math/rand"
)

// QuickSort is a divide and conquer algorithm. It partitions the array first
// and combines after.
func QuickSort(arr []int) []int {
	if !(len(arr) > 1) {
		return arr
	}

	// Choose random pivot
	i := rand.Intn(len(arr))
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	p := partition(arr, 0, len(arr)-1)
	QuickSort(arr[0:p])
	QuickSort(arr[p+1:])

	return arr
}

// Lomuto's partitioning scheme
func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

/*

	Trivial implementation

func QuickSort(arr []int) []int {
	if !(len(arr) > 1) {
		return arr
	}

	pivot := len(arr) / 2

	less := make([]int, 0)
	equal := make([]int, 0)
	greater := make([]int, 0)

	for _, element := range arr {
		if element < arr[pivot] {
			less = append(less, element)
		} else if element > arr[pivot] {
			greater = append(greater, element)
		} else {
			equal = append(equal, element)
		}
	}

	less = QuickSort(less)
	greater = QuickSort(greater)

	return append(less, append(equal, greater...)...)
}

*/
