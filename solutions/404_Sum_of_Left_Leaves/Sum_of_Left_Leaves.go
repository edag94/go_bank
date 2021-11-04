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
		if isLeaf(root.Left) {
			sum += root.Left.Val
		}
		sum += sumOfLeftLeaves(root.Left)
		sum += sumOfLeftLeaves(root.Right)
	}
	return sum
}
