package searchAlgorithms

import "fmt"

// TreeNode defines a node in the tree
type TreeNode struct {
	val    int
	index  int
	left   *TreeNode
	right  *TreeNode
	height int
}

// insert inserts a value into the tree
func insert(root *TreeNode, value int, index int) *TreeNode {
	if root == nil {
		return &TreeNode{
			val:   value,
			index: index,
		}
	}
	if value < root.val {
		root.left = insert(root.left, value, index)
	} else {
		root.right = insert(root.right, value, index)
	}
	return root
}

// search finds target value's index
func search(root *TreeNode, target int) int {
	if root == nil {
		return -1
	}
	if root.val == target {
		return root.index
	}
	if target < root.val {
		return search(root.left, target)
	} else {
		return search(root.right, target)
	}
}

// delete removes target value's node
func remove(root *TreeNode, target int) *TreeNode {
	// no value, return nil
	if root == nil {
		return nil
	}

	// try to find the target node
	if target < root.val {
		root.left = remove(root.left, target)
	} else if target > root.val {
		root.right = remove(root.right, target)
	} else {
		// case 1: node has no children leaf node)
		if root.left == nil && root.right == nil {
			// remove it
			root = nil
		} else if root.left == nil {
			// case 2: node has one child
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			// case 3: node has two children
			// find min from right or find max from left as successor
			successor := findMin(root.right)
			root.val = successor.val
			root.index = successor.index
			// Delete the inorder successor
			root.right = remove(root.right, successor.val)
		}
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	current := root
	for current != nil && current.left != nil {
		current = current.left
	}
	return current
}

// inorderTraversal performs an inorder traversal of the tree
func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.left)  // traverse the left subtree
		fmt.Print(root.val, " ")     // print the node's value
		inorderTraversal(root.right) // traverse the right subtree
	}
}

func BinarySearchTree(arr []int, target int) int {

	arr = []int{10, 9, 3, 2, 4}
	target = 3
	// set up the root node of the tree
	var root *TreeNode
	// build the tree
	for index, value := range arr {
		root = insert(root, value, index)
	}
	// print the tree in ascending order
	//inorderTraversal(root)

	// find target number from the tree
	index := search(root, target)
	newTree := remove(root, target)
	if newTree != nil {
		inorderTraversal(newTree)
	} else {
		fmt.Println("cannot find the target node!")
	}

	return index
}
