package main

import (
	"fmt"
	"practice/internal/algorithms/searchAlgorithms"
	"practice/internal/pkg/number"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := searchAlgorithms.AVLTreeSearch(number.DisOrderedNumbers, 133)
	fmt.Println("result is : ", r)

}
