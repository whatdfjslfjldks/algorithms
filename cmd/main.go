package main

import (
	"fmt"
	"practice/internal/algorithms/searchAlgorithms"
	"practice/internal/pkg/number"
)

func main() {
	r := searchAlgorithms.BlockSearch(number.Numbers, 150)
	fmt.Println("dsf:", r)
}
