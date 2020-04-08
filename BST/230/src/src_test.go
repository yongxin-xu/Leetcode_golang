package src_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, k *int, ret *int) {
	if root == nil {
		return
	}
	helper(root.Left, k, ret)
	if *k == 0 {
		return
	}
	(*k)--
	if *k == 0 {
		*ret = root.Val
		return
	}
	helper(root.Right, k, ret)
}

func kthSmallest(root *TreeNode, k int) int {
	ret := 0
	helper(root, &k, &ret)
	return ret
}

func TestFn(t *testing.T) {
	r := new(TreeNode)
	r.Val = 5
	r.Left = new(TreeNode)
	r.Left.Val = 3
	r.Right = new(TreeNode)
	r.Right.Val = 6
	r.Left.Left = new(TreeNode)
	r.Left.Left.Val = 2
	r.Left.Right = new(TreeNode)
	r.Left.Right.Val = 4
	r.Left.Left.Left = new(TreeNode)
	r.Left.Left.Left.Val = 1
	ret := kthSmallest(r, 3)
	t.Log(ret)
}
