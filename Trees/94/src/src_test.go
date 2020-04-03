package src_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, ret *[]int) {
	if root.Left != nil {
		helper(root.Left, ret)
	}

	*ret = append(*ret, root.Val)

	if root.Right != nil {
		helper(root.Right, ret)
	}
}

func inorderTraversalRecursive(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	helper(root, &ret)
	return ret
}

func inorderTraversalIterative(root *TreeNode) []int {
	ret := []int{}
	stack := [](*TreeNode){}
	ptr := root

	for ptr != nil || len(stack) != 0 {
		for ptr != nil {
			stack = append(stack, ptr) // stack.Push(ptr)
			ptr = ptr.Left
		}

		if len(stack) != 0 {
			ptr = stack[len(stack)-1] // stack.Top()
			ret = append(ret, ptr.Val)
			stack = stack[0 : len(stack)-1] // stack.Pop()
			ptr = ptr.Right
		}
	}

	return ret
}

func inorderTraversalMorris(root *TreeNode) []int {
	ret := []int{}

	ptr := root

	for ptr != nil {
		if ptr.Left == nil {
			ret = append(ret, ptr.Val)
			ptr = ptr.Right
		} else {
			prev := ptr.Left
			for prev.Right != nil && prev.Right != ptr {
				prev = prev.Right
			}

			if prev.Right == nil {
				prev.Right = ptr
				ptr = ptr.Left
			} else {
				ret = append(ret, ptr.Val)
				prev.Right = nil
				ptr = ptr.Right
			}
		}
	}

	return ret
}

func TestExample1(t *testing.T) {
	t1 := new(TreeNode)
	t1.Val = 5
	t1.Left = new(TreeNode)
	t1.Left.Val = 2
	t1.Left.Left = new(TreeNode)
	t1.Left.Left.Val = 1
	t1.Left.Right = new(TreeNode)
	t1.Left.Right.Val = 4
	t1.Left.Right.Left = new(TreeNode)
	t1.Left.Right.Left.Val = 3
	t1.Right = new(TreeNode)
	t1.Right.Val = 8
	t1.Right.Left = new(TreeNode)
	t1.Right.Left.Val = 6
	t1.Right.Left.Right = new(TreeNode)
	t1.Right.Left.Right.Val = 7
	t.Log(inorderTraversalRecursive(t1))
	t.Log(inorderTraversalIterative(t1))
	t.Log(inorderTraversalMorris(t1))
}
