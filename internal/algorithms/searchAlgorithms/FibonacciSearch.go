package searchAlgorithms

import (
	"math"
	"sort"
)

// FibonacciSearch performs a Fibonacci search on a sorted array
func FibonacciSearch(arr []int, target int) int {
	// sort first
	sort.Ints(arr)

	n := len(arr)

	// Initialize Fibonacci numbers
	fibM2 := 0           // (m-2)'th Fibonacci number
	fibM1 := 1           // (m-1)'th Fibonacci number
	fib := fibM2 + fibM1 // m'th Fibonacci number

	// Find the smallest Fibonacci number greater than or equal to n
	for fib < n {
		fibM2 = fibM1
		fibM1 = fib
		fib = fibM2 + fibM1
	}

	// Mark the eliminated range from the front
	offset := -1

	// While there are elements to be checked
	for fib > 1 {
		// Find the index to be checked
		i := int(math.Min(float64(offset+fibM2), float64(n-1)))

		// If target is found
		if arr[i] == target {
			return i
		}

		// If target is greater, ignore the left half
		if arr[i] < target {
			fib = fibM1
			fibM1 = fibM2
			fibM2 = fib - fibM1
			offset = i
		} else {
			// If target is smaller, ignore the right half
			fib = fibM2
			fibM1 = fibM1 - fibM2
			fibM2 = fib - fibM1
		}
	}

	// Compare the last element with target
	if fibM1 == 1 && arr[offset+1] == target {
		return offset + 1
	}

	// If the element is not found, return -1
	return -1
}
