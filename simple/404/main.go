// 左叶子之和

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//// 深度优先
//func sumOfLeftLeaves(root *TreeNode) int {
//	sum := 0
//	if root == nil {
//		return 0
//	}
//	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
//		sum += root.Left.Val
//	}
//	sum += sumOfLeftLeaves(root.Left)
//	sum += sumOfLeftLeaves(root.Right)
//	return sum
//}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

// 广度优先
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return 0
	}
	rs := []*TreeNode{root}
	for len(rs) > 0 {
		node := rs[0]
		rs = rs[1:]
		if node.Left != nil {
			if isLeafNode(node.Left) {
				sum += node.Left.Val
			} else {
				rs = append(rs, node.Left)
			}
		}
		if node.Right != nil && !isLeafNode(node.Right) {
			rs = append(rs, node.Right)
		}
	}
	return sum
}

func main() {
	a := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(sumOfLeftLeaves(a))
}
