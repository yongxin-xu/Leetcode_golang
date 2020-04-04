package src_test

import "testing"

func moveZeroes(nums []int) {
	zero_tail := -1
	for i, num := range nums {
		if num == 0 {
			if zero_tail == -1 {
				zero_tail = i
			}
		} else {
			if zero_tail != -1 {
				temp := zero_tail
				zero_tail++
				nums[temp], nums[i] = nums[i], nums[temp]
			}
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	a := []int{2, 0, 3, 0, 0, 1, 0, 3, 12, 0}
	moveZeroes(a)
	t.Log(a)
}
