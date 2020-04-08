package src_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	// Two pointers, fastPtr and slowPtr
	// fastPtr moves two steps per iteration, slowPtr moves one step
	fastPtr := head
	slowPtr := head

	for fastPtr != nil {
		fastPtr = fastPtr.Next
		if fastPtr == nil {
			break
		}
		fastPtr = fastPtr.Next
		slowPtr = slowPtr.Next
	}
	return slowPtr
}
