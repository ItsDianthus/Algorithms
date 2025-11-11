package main

import "fmt"

// easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.Val + SumBST(root.Right) + SumBST(root.Left)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return rangeSumBST(root.Left, low, high)
}

func main() {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 18}

	// Тестируем сумму в диапазоне [7, 15]
	// low, high := 7, 15
	result := SumBST(root)

	fmt.Printf("Сумма значений в диапазоне [ %d ]:", result)

}
