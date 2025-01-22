package searchAlgorithms

import "sort"

func InterpolationSearch(arr []int, target int) int {
	// same as binarySearch, sort the arr in ascending order firstly,
	sort.Ints(arr)

	left := 0
	right := len(arr) - 1
	for left <= right {
		if arr[left] == arr[right] {
			if arr[left] == target {
				return left
			}
			break
		}
		//mid := left + ((target-arr[left])/(arr[right]-arr[left]))*(right-left)
		mid := left + (target-arr[left])*((right-left)/(arr[right]-arr[left]))

		if target > arr[mid] {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		} else if target == arr[mid] {
			// return the index
			return mid
		}
	}
	return -1
}
