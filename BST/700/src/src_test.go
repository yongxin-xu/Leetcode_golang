package src_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val == val {
			return root
		} else if val < root.Val {
			return searchBST(root.Left, val)
		} else {
			return searchBST(root.Right, val)
		}
	}
	return nil
}

func TestFunction(t *testing.T) {
	root := new(TreeNode)
	root.Val = 4
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 1
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	root.Right = new(TreeNode)
	root.Right.Val = 7
	t.Log(searchBST(root, 2))
}
