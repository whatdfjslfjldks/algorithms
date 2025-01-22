package searchAlgorithms

import (
	"fmt"
	"math"
	"sort"
)

func JumpSearch(arr []int, target int) int {
	// sort the arr in ascending order first
	sort.Ints(arr)

	n := len(arr)
	blockSize := int(math.Sqrt(float64(n)))

	for i := 0; i < n; i += blockSize {
		end := int(math.Min(float64(i+blockSize), float64(n)))

		fmt.Println("Checking block:", i, "to", end)

		for j := i; j < end; j++ {
			if arr[j] == target {
				return j
			}
		}
		// If target is larger than the current block's last element
		if arr[i] > target {
			break
		}
	}
	return -1
}
