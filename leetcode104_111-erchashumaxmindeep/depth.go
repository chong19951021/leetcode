package leetcode104_111_erchashumaxmindeep

import "coding/common"

func maxDepth(root *common.TreeNode) int {

	var depth int
	return dfs(root, depth)

}

func dfs(node *common.TreeNode, depthCur int) int {
	if node == nil {
		return depthCur
	}
	depthCur++
	return common.Max(dfs(node.Left, depthCur), dfs(node.Right, depthCur))

}

func MinDepth(root *common.TreeNode) int {
	var depth int
	dfsMin(root, depth)
	return res

}

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

var res = 0
