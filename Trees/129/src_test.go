package src_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum *int, cur int) {
	if root == nil {
		return
	}
	cur = cur*10 + root.Val
	if root.Right == nil && root.Left == nil {
		*sum += cur
		return
	}
	dfs(root.Left, sum, cur)
	dfs(root.Right, sum, cur)
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	dfs(root, &sum, 0)
	return sum
}

func TestFunction(t *testing.T) {
	r := new(TreeNode)
	r.Val = 4
	rL := new(TreeNode)
	rL.Val = 9
	rR := new(TreeNode)
	rR.Val = 0
	rLL := new(TreeNode)
	rLL.Val = 5
	rLR := new(TreeNode)
	rLR.Val = 1
	r.Left = rL
	r.Right = rR
	r.Left.Left = rLL
	r.Left.Right = rLR
	t.Log(sumNumbers(r))
}
