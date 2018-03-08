package search

// BinarySearch searches for target element in given sorted array.
func BinarySearch(arr []int, key int) int {
	lower := 0
	upper := len(arr) - 1
	for lower <= upper {
		mid := lower + (upper-lower)/2
		if arr[mid] > key {
			upper = mid - 1
		} else if arr[mid] < key {
			lower = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
