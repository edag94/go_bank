package Sum_of_Left_Leaves

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
