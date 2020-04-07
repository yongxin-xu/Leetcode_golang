package src_test

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode = nil

func inOrderTraversal(root *TreeNode, minDiff *int) {
	if root != nil {
		inOrderTraversal(root.Left, minDiff)
		if prev != nil && (root.Val-prev.Val < *minDiff) {
			*minDiff = root.Val - prev.Val
		}
		prev = root
		inOrderTraversal(root.Right, minDiff)
	}
}

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt32
	prev = nil
	// Since this is a Binary Search Tree,
	// an inorder traversal would bring a sorted array
	inOrderTraversal(root, &minDiff)
	return minDiff
}

func TestFunction(t *testing.T) {
	root := new(TreeNode)
	root.Right = new(TreeNode)
	root.Right.Left = new(TreeNode)
	root.Val = 1
	root.Right.Val = 5
	root.Right.Left.Val = 3
	t.Log(getMinimumDifference(root))
}
