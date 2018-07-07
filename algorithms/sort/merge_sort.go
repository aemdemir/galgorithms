package sort

// MergeSort is a divide and conquer algorithm. It splits the array first
// and merge after.
func MergeSort(array []int) []int {
	if !(len(array) > 1) {
		return array
	}

	middle := len(array) / 2

	leftArr := MergeSort(array[:middle])
	rightArr := MergeSort(array[middle:])

	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	mergedArr := make([]int, 0)

	i, j := 0, 0

	for i < len(leftArr) && j < len(rightArr) {
		switch {
		case leftArr[i] < rightArr[j]:
			mergedArr = append(mergedArr, leftArr[i])
			i++
		case rightArr[j] < leftArr[i]:
			mergedArr = append(mergedArr, rightArr[j])
			j++
		default:
			mergedArr = append(mergedArr, leftArr[i], rightArr[j])
			i++
			j++
		}
	}

	if i < len(leftArr) {
		mergedArr = append(mergedArr, leftArr[i:]...)
	}

	if j < len(rightArr) {
		mergedArr = append(mergedArr, rightArr[j:]...)
	}

	return mergedArr
}
