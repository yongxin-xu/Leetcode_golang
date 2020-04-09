package src_test

import (
	"sort"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getInorderArray(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	getInorderArray(root.Left, arr)
	*arr = append(*arr, root.Val)
	getInorderArray(root.Right, arr)
}

func recoverTreeImpl(root *TreeNode, arr *[]int, idx *int) {
	if root == nil {
		return
	}
	recoverTreeImpl(root.Left, arr, idx)
	root.Val = (*arr)[*idx]
	(*idx)++
	recoverTreeImpl(root.Right, arr, idx)
}

func recoverTree(root *TreeNode) {
	arr := []int{}
	idx := 0
	getInorderArray(root, &arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	recoverTreeImpl(root, &arr, &idx)
}

func TestFn(t *testing.T) {
	r := new(TreeNode)
	r.Val = 1
	r.Left = new(TreeNode)
	r.Left.Val = 3
	r.Left.Right = new(TreeNode)
	r.Left.Right.Val = 2
	recoverTree(r)
}
