package leetcode114_erchashuzhankaiweilianbiao

import "coding/common"

func flatten(root *common.TreeNode) {
	for root != nil {
		nextNode := root.Right
		oldLeft := root.Left
		if root.Left != nil {
			var lastNode *common.TreeNode
			nextRightNode := root.Left.Right
			if nextRightNode != nil {
				for nextRightNode.Right != nil {
					nextRightNode = nextRightNode.Right
				}
				lastNode = nextRightNode
			} else {
				lastNode = root.Left
			}
			lastNode.Right = nextNode
			root.Left = nil
			root.Right = oldLeft
		}
		root = root.Right
	}

}
