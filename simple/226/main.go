//  翻转二叉树

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 翻转二叉树只要从上至下左右互换就可以了
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil && root.Left == nil {
		return root
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Right)
	invertTree(root.Left)
	return root
}
