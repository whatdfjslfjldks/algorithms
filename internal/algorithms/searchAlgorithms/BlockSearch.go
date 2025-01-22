package searchAlgorithms

import (
	"math"
	"sort"
)

func BlockSearch(arr []int, target int) int {
	// same, work for sorted arr
	sort.Ints(arr)

	n := len(arr)
	blockSize := int(math.Sqrt(float64(n)))

	start := 0
	end := blockSize

	for start < n {
		// find the target block
		if arr[start] <= target && target <= arr[min(end, n)-1] {
			break
		}
		start = end
		end += blockSize
		if end > n {
			end = n
		}
	}

	if start >= n {
		return -1
	}

	for i := start; i < end; i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
