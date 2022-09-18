package main

import (
	"coding/common"
	"fmt"
)

// 二叉树最短最长深度

func main() {
	root := &common.TreeNode{
		Val:  2,
		Left: nil,
		Right: &common.TreeNode{
			Val:  3,
			Left: nil,
			Right: &common.TreeNode{
				Val:  4,
				Left: nil,
				Right: &common.TreeNode{
					Val:  5,
					Left: nil,
					Right: &common.TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	depth := minDepth(root)
	fmt.Println(depth)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *common.TreeNode) int {
	var depth int
	dfsMin(root, depth)
	return res

}

var res = 0

func dfsMin(node *common.TreeNode, depthCur int) {
	if node == nil {
		return
	}
	depthCur++
	if node.Left == nil && node.Right == nil {
		if depthCur < res || res == 0 {
			res = depthCur
		}
	}

	dfsMin(node.Left, depthCur)
	dfsMin(node.Right, depthCur)

}
