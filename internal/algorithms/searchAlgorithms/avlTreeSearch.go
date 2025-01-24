package searchAlgorithms

import "fmt"

// avlInsert 插入元素并确保树的平衡
func avlInsert(root *TreeNode, value int, index int) *TreeNode {
	// 插入新节点
	if root == nil {
		return &TreeNode{
			val:    value,
			index:  index,
			height: 1, // 新插入节点的高度为1
		}
	}

	if value < root.val {
		root.left = avlInsert(root.left, value, index)
	} else if value > root.val {
		root.right = avlInsert(root.right, value, index)
	} else {
		// 如果相同的值，直接返回，不插入
		return root
	}

	// 更新当前节点的高度
	root.height = 1 + max(calHeight(root.left), calHeight(root.right))
	fmt.Println(root.height)

	// 检查平衡因子并执行旋转
	balanceFactor := calBalanceFactor(root)

	// 左左情况
	if balanceFactor > 1 && value < root.left.val {
		return rotateRight(root)
	}

	// 右右情况
	if balanceFactor < -1 && value > root.right.val {
		return rotateLeft(root)
	}

	// 左右情况
	if balanceFactor > 1 && value > root.left.val {
		root.left = rotateLeft(root.left)
		return rotateRight(root)
	}

	// 右左情况
	if balanceFactor < -1 && value < root.right.val {
		root.right = rotateRight(root.right)
		return rotateLeft(root)
	}

	return root
}

// calHeight 计算节点的高度
func calHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.height
}

// calBalanceFactor 计算节点的平衡因子
func calBalanceFactor(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return calHeight(root.left) - calHeight(root.right)
}

// rotateRight 右旋操作
func rotateRight(y *TreeNode) *TreeNode {
	x := y.left
	T2 := x.right

	// 执行旋转
	x.right = y
	y.left = T2

	// 更新节点的高度
	y.height = max(calHeight(y.left), calHeight(y.right)) + 1
	x.height = max(calHeight(x.left), calHeight(x.right)) + 1

	return x
}

// rotateLeft 左旋操作
func rotateLeft(x *TreeNode) *TreeNode {
	y := x.right
	T2 := y.left

	// 执行旋转
	y.left = x
	x.right = T2

	// 更新节点的高度
	x.height = max(calHeight(x.left), calHeight(x.right)) + 1
	y.height = max(calHeight(y.left), calHeight(y.right)) + 1

	return y
}

// max 返回两个数中较大的值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// AVLTreeSearch 使用 AVL 树搜索
func AVLTreeSearch(arr []int, target int) int {

	arr = []int{10, 9, 3, 2, 4, 5}
	target = 3
	var root *TreeNode

	for index, value := range arr {
		root = avlInsert(root, value, index)
	}

	return -1
}
