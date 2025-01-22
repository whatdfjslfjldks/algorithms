package searchAlgorithms

// HashSearch Does not cover methods to resolve hash collisions
func HashSearch(arr []int, target int) int {
	hashTable := make(map[int]int)

	for i, val := range arr {
		hashTable[val] = i
	}

	if index, found := hashTable[target]; found {
		return index
	}
	return -1
}
