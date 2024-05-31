//  二叉树的所有路径

package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
//func binaryTreePaths(root *TreeNode) []string {
//	var result []string
//	if root == nil {
//		return result
//	}
//	var dfs func(node *TreeNode, path *[]string, r *[]string)
//	dfs = func(r *TreeNode, path *[]string, res *[]string) {
//		if r == nil {
//			return
//		}
//		*path = append(*path, strconv.Itoa(r.Val))
//		if r.Right == nil && r.Left == nil {
//			*res = append(*res, strings.Join(*path, "->"))
//		}
//		dfs(r.Left, path, res)
//		dfs(r.Right, path, res)
//		if len(*path) > 0 {
//			*path = (*path)[:len(*path)-1]
//		}
//	}
//	var path []string
//	dfs(root, &path, &result)
//	return result
//}

// BFS
func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	nodeQueue := []*TreeNode{}
	pathQueue := []string{}
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}
	fmt.Println(binaryTreePaths(root))
}
