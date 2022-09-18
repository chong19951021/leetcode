package leetcode94_145_144_bianlierchashu

import "coding/common"

// 前序遍历
func preorderTraversal(root *common.TreeNode) []int {
	res := make([]int, 0)
	var fun1 func(root *common.TreeNode)
	fun1 = func(root *common.TreeNode) {
		if root == nil {
			return

		}
		res = append(res, root.Val)
		fun1(root.Left)
		fun1(root.Right)

	}
	fun1(root)
	return res

}

// 后序遍历
func postorderTraversal(root *common.TreeNode) []int {

	res := make([]int, 0)
	var fun1 func(root *common.TreeNode)
	fun1 = func(root *common.TreeNode) {
		if root == nil {
			return

		}
		fun1(root.Left)
		fun1(root.Right)
		res = append(res, root.Val)

	}
	fun1(root)
	return res

}

// 中序遍历
func inorderTraversal(root *common.TreeNode) []int {
	res := make([]int, 0)
	var fun1 func(root *common.TreeNode)
	fun1 = func(root *common.TreeNode) {
		if root == nil {
			return

		}
		fun1(root.Left)
		res = append(res, root.Val)

		fun1(root.Right)

	}
	fun1(root)
	return res
}
