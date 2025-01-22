package searchAlgorithms

import "sort"

func BinarySearch(arr []int, target int) int {
	// first of all, sort the array in ascending order
	sort.Ints(arr)

	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if target > arr[mid] {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		} else if target == arr[mid] {
			// return index
			return mid
		}
	}
	return -1
}
