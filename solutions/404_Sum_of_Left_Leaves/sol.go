package Sum_of_Left_Leaves

import (
	"github.com/edag94/go_bank/utils"
)

type TreeNode = utils.TreeNode

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum int = 0
	if root == nil {
		return sum
	} else {
		// if left node is leaf, add to sum
		if isLeaf(root.Left) {
			sum += root.Left.Val
		}

		// recurse further down tree
		sum += sumOfLeftLeaves(root.Left)
		sum += sumOfLeftLeaves(root.Right)
	}
	return sum
}

func sumOfLeftLeavesClosure(root *TreeNode) int {
	var sum int = 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root != nil {
			// if left node is leaf, add to sum
			if isLeaf(root.Left) {
				sum += root.Left.Val
			}

			// recurse further down tree
			if root.Left != nil {
				helper(root.Left)
			}
			if root.Right != nil {
				helper(root.Right)
			}
		}
	}
	helper(root)
	return sum
}
