package src_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, max *TreeNode, min *TreeNode) bool {
	if root == nil {
		return true
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	if min != nil && root.Val <= min.Val {
		return false
	}

	return helper(root.Left, root, min) &&
		helper(root.Right, max, root)
}

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func TestBST(t *testing.T) {

}
