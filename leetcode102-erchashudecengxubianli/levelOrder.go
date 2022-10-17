package leetcode102_erchashudecengxubianli

import "coding/common"

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
func levelOrder(root *common.TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*common.TreeNode{root}
	for len(queue) > 0 {
		var nextQueue []*common.TreeNode
		curQueue := []int{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
			curQueue = append(curQueue, queue[i].Val)
		}
		res = append(res, curQueue)
		queue = nextQueue

	}
	return res
}
