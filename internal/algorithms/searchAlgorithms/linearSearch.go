package searchAlgorithms

func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			// return index
			return i
		}
	}
	return -1
}
