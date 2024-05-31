//  完全二叉树的节点个数

package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//	count := 1
//	count += countNodes(root.Left)
//	count += countNodes(root.Right)
//	return count
//}

// 二分查找结合位运算
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			// 表示左节点会变为中间值
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		// 最终结果为false,表示(k + 1)会变为左侧值
		// 最终结果为true,表示k会变为右侧值
		return node == nil
	}) - 1
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}
	fmt.Println(countNodes(root))
}
