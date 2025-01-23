package main

import (
	"fmt"
	"practice/internal/algorithms/searchAlgorithms"
	"practice/internal/pkg/number"
)

func main() {
	r := searchAlgorithms.BinarySearchTree(number.DisOrderedNumbers, 133)
	fmt.Println("result is : ", r)
}
